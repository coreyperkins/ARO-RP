# Prepare a shared RP development environment

Follow these steps to build a shared RP development environment and secrets
file.  A single RP development environment can be shared across multiple
developers and/or CI flows.  It may include multiple resource groups in multiple
locations.

## Prerequisites

1. You will need `Contributor` and `User Access Administrator` roles on your
   Azure subscription, as well as the ability to create and configure AAD
   applications.

1. You will need a publicly resolvable DNS Zone resource in your Azure
   subscription.  Set PARENT_DOMAIN_NAME and PARENT_DOMAIN_RESOURCEGROUP to the name and
   resource group of the DNS Zone resource:

   ```bash
   PARENT_DOMAIN_NAME=osadev.cloud
   PARENT_DOMAIN_RESOURCEGROUP=dns
   ```

1. You will need a storage account in your Azure subscription in which to store
   shared development environment secrets.   The storage account must contain a
   private container named `secrets`.  All team members must have `Storage Blob
   Data Reader` or `Storage Blob Data Contributor` role on the storage account.
   Set SECRET_SA_ACCOUNT_NAME to the name of the storage account:

   ```bash
   SECRET_SA_ACCOUNT_NAME=rharosecrets
   ```

1. You will need an AAD object (this could be your AAD user, or an AAD group of
   which you are a member) which will be able to administer certificates in the
   development environment key vault(s).  Set ADMIN_OBJECT_ID to the object ID.

   ```bash
   ADMIN_OBJECT_ID="$(az ad group show -g Engineering --query objectId -o tsv)"
   ```

1. You will need the ARO RP-specific pull secret (ask one of the
   @azure-red-hat-openshift GitHub team for this):

   ```bash
   PULL_SECRET=...
   ```

1. Install [Go 1.14](https://golang.org/dl) or later, if you haven't already.

1. Install the [Azure
   CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli), if you
   haven't already.

1. Log in to Azure:

   ```bash
   az login

   AZURE_TENANT_ID=$(az account show --query tenantId -o tsv)
   AZURE_SUBSCRIPTION_ID=$(az account show --query id -o tsv)
   ```

1. Git clone this repository to your local machine:

   ```bash
   go get -u github.com/Azure/ARO-RP/...
   cd ${GOPATH:-$HOME/go}/src/github.com/Azure/ARO-RP
   ```

1. Prepare the secrets directory:

   ```bash
   mkdir -p secrets
   ```


## AAD applications

1. Create an AAD application which will fake up the ARM layer:

   ```bash
   AZURE_ARM_CLIENT_SECRET="$(uuidgen)"
   AZURE_ARM_CLIENT_ID="$(az ad app create \
     --display-name aro-v4-arm-shared \
     --end-date '2299-12-31T11:59:59+00:00' \
     --identifier-uris "https://$(uuidgen)/" \
     --key-type password \
     --password "$AZURE_ARM_CLIENT_SECRET" \
     --query appId \
     -o tsv)"
   az ad sp create --id "$AZURE_ARM_CLIENT_ID" >/dev/null
   ```

   Later this application will be granted:

   * `User Access Administrator` on your subscription.

1. Create an AAD application which will fake up the first party application.

   This application requires client certificate authentication to be enabled.  A
   suitable key/certificate file can be generated using the following helper
   utility:

   ```bash
   go run ./hack/genkey -client firstparty
   mv firstparty.* secrets
   ```

   Now create the application:

   ```bash
   AZURE_FP_CLIENT_ID="$(az ad app create \
     --display-name aro-v4-fp-shared \
     --identifier-uris "https://$(uuidgen)/" \
     --query appId \
     -o tsv)"
   az ad app credential reset \
     --id "$AZURE_FP_CLIENT_ID" \
     --cert "$(base64 -w0 <secrets/firstparty.crt)" >/dev/null
   az ad sp create --id "$AZURE_FP_CLIENT_ID" >/dev/null
   ```

   Later this application will be granted:

   * `ARO v4 FP Subscription` on your subscription.
   * `DNS Zone Contributor` on the DNS zone in RESOURCEGROUP.
   * `Network Contributor` on RESOURCEGROUP.

1. Create an AAD application which will fake up the RP identity.

   ```bash
   AZURE_RP_CLIENT_SECRET="$(uuidgen)"
   AZURE_RP_CLIENT_ID="$(az ad app create \
     --display-name aro-v4-rp-shared \
     --end-date '2299-12-31T11:59:59+00:00' \
     --identifier-uris "https://$(uuidgen)/" \
     --key-type password \
     --password "$AZURE_RP_CLIENT_SECRET" \
     --query appId \
     -o tsv)"
   az ad sp create --id "$AZURE_RP_CLIENT_ID" >/dev/null
   ```

   Later this application will be granted:

   * `Reader` on RESOURCEGROUP.
   * `Secrets / Get` on the key vault in RESOURCEGROUP.
   * `DocumentDB Account Contributor` on the CosmosDB resource in RESOURCEGROUP.

1. Create an AAD application which will be used by E2E and tooling.

   ```bash
   AZURE_CLIENT_SECRET="$(uuidgen)"
   AZURE_CLIENT_ID="$(az ad app create \
     --display-name aro-v4-tooling-shared \
     --end-date '2299-12-31T11:59:59+00:00' \
     --identifier-uris "https://$(uuidgen)/" \
     --key-type password \
     --password "$AZURE_CLIENT_SECRET" \
     --query appId \
     -o tsv)"
   az ad sp create --id "$AZURE_CLIENT_ID" >/dev/null
   ```

   Later this application will be granted:

   * `Contributor`  on your subscription.
   * `User Access Administrator` on your subscription.

   You must also manually grant this application the `Microsoft.Graph/Application.ReadWrite.OwnedBy` permission, which requires admin access, in order for AAD applications to be created/deleted on a per-cluster basis.

   * Go into the Azure Portal
   * Go to Azure Active Directory
   * Navigate to the `aro-v4-tooling-shared` app page
   * Click 'API permissions' in the left side pane
   * Click 'Microsoft Graph'
   * Click 'Add a permission'.
   * Select 'Application permissions'
   * Search for 'Application' and select `Application.ReadWrite.OwnedBy`
   * Click 'Add permissions'
   * This request will need to be approved by a tenant administrator. If you are one, you can click the `Grant admin consent for <name>` button to the right of the `Add a permission` button on the app page

1. Set up the RP role definitions and subscription role assignments in your
   Azure subscription. This mimics the RBAC that ARM sets up.  With at least
   `User Access Administrator` permissions on your subscription, do:

   ```bash
   az deployment sub create \
     -l $LOCATION \
     --template-file deploy/rbac-development.json \
     --parameters \
       "armServicePrincipalId=$(az ad sp list --filter "appId eq '$AZURE_ARM_CLIENT_ID'" --query '[].objectId' -o tsv)" \
       "fpServicePrincipalId=$(az ad sp list --filter "appId eq '$AZURE_FP_CLIENT_ID'" --query '[].objectId' -o tsv)" \
       "devServicePrincipalId=$(az ad sp list --filter "appId eq '$AZURE_CLIENT_ID'" --query '[].objectId' -o tsv)" \
     >/dev/null
   ```

1. Create an AAD application which will fake up the portal client.

   This application requires client certificate authentication to be enabled.  A
   suitable key/certificate file can be generated using the following helper
   utility:

   ```bash
   go run ./hack/genkey -client portal-client
   mv portal-client.* secrets
   ```

   ```bash
   AZURE_PORTAL_CLIENT_ID="$(az ad app create \
     --display-name aro-v4-portal-shared \
     --identifier-uris "https://$(uuidgen)/" \
     --reply-urls "https://localhost:8444/callback" \
     --query appId \
     -o tsv)"
   az ad app credential reset \
     --id "$AZURE_PORTAL_CLIENT_ID" \
     --cert "$(base64 -w0 <secrets/portal-client.crt)" >/dev/null
   ```

   TODO: more steps are needed to configure aro-v4-portal-shared.


## Certificates

1. Create the VPN CA key/certificate.  A suitable key/certificate file can be
   generated using the following helper utility:

   ```bash
   go run ./hack/genkey -ca vpn-ca
   mv vpn-ca.* secrets
   ```

1. Create the VPN client key/certificate.  A suitable key/certificate file can be
   generated using the following helper utility:

   ```bash
   go run ./hack/genkey -client -keyFile secrets/vpn-ca.key -certFile secrets/vpn-ca.crt vpn-client
   mv vpn-client.* secrets
   ```

1. Create the proxy serving key/certificate.  A suitable key/certificate file
   can be generated using the following helper utility:

   ```bash
   go run ./hack/genkey proxy
   mv proxy.* secrets
   ```

1. Create the proxy client key/certificate.  A suitable key/certificate file can
   be generated using the following helper utility:

   ```bash
   go run ./hack/genkey -client proxy-client
   mv proxy-client.* secrets
   ```

1. Create the proxy ssh key/certificate.  A suitable key/certificate file can
   be generated using the following helper utility:

   ```bash
   ssh-keygen -f secrets/proxy_id_rsa -N ''
   ```

1. Create an RP serving key/certificate.  A suitable key/certificate file
   can be generated using the following helper utility:

   ```bash
   go run ./hack/genkey localhost
   mv localhost.* secrets
   ```


# Environment file

1. Choose the resource group prefix.  The resource group location will be
   appended to the prefix to make the resource group name.

   ```bash
   RESOURCEGROUP_PREFIX=v4
   ```

1. Choose the proxy domain name label.  This final proxy hostname will be of the
   form `vm0.$PROXY_DOMAIN_NAME_LABEL.$LOCATION.cloudapp.azure.com`.

   ```bash
   PROXY_DOMAIN_NAME_LABEL=aroproxy
   ```

1. Create the secrets/env file:

   ```bash
   cat >secrets/env <<EOF
   export AZURE_TENANT_ID='$AZURE_TENANT_ID'
   export AZURE_SUBSCRIPTION_ID='$AZURE_SUBSCRIPTION_ID'
   export AZURE_ARM_CLIENT_ID='$AZURE_ARM_CLIENT_ID'
   export AZURE_ARM_CLIENT_SECRET='$AZURE_ARM_CLIENT_SECRET'
   export AZURE_FP_CLIENT_ID='$AZURE_FP_CLIENT_ID'
   export AZURE_PORTAL_CLIENT_ID='$AZURE_PORTAL_CLIENT_ID'
   export AZURE_PORTAL_ACCESS_GROUP_IDS='$ADMIN_OBJECT_ID'
   export AZURE_PORTAL_ELEVATED_GROUP_IDS='$ADMIN_OBJECT_ID'
   export AZURE_CLIENT_ID='$AZURE_CLIENT_ID'
   export AZURE_CLIENT_SECRET='$AZURE_CLIENT_SECRET'
   export AZURE_RP_CLIENT_ID='$AZURE_RP_CLIENT_ID'
   export AZURE_RP_CLIENT_SECRET='$AZURE_RP_CLIENT_SECRET'
   export RESOURCEGROUP="$RESOURCEGROUP_PREFIX-\$LOCATION"
   export PROXY_HOSTNAME="vm0.$PROXY_DOMAIN_NAME_LABEL.\$LOCATION.cloudapp.azure.com"
   export DATABASE_NAME="\$USER"
   export RP_MODE='development'
   export PULL_SECRET='$PULL_SECRET'
   export SECRET_SA_ACCOUNT_NAME='$SECRET_SA_ACCOUNT_NAME'
   export DATABASE_ACCOUNT_NAME="\$RESOURCEGROUP"
   export KEYVAULT_PREFIX="\$RESOURCEGROUP"
   ADMIN_OBJECT_ID='$ADMIN_OBJECT_ID'
   PARENT_DOMAIN_NAME='$PARENT_DOMAIN_NAME'
   PARENT_DOMAIN_RESOURCEGROUP='$PARENT_DOMAIN_RESOURCEGROUP'
   export DOMAIN_NAME="\$LOCATION.\$PARENT_DOMAIN_NAME"
   export AZURE_ENVIRONMENT='AzurePublicCloud'
   EOF
   ```

After creating this file, run `make secrets-update` to upload it to your
storage account so other people on your team can access it via `make secrets`.

## Deploy shared RP development environment (once per location)

Look at the [helper file](../hack/devtools/deploy-shared-env.sh) to understand
each of the bash functions below.

1. Copy, edit (if necessary) and source your environment file.  The required
   environment variable configuration is documented immediately below:

   ```bash
   cp env.example env
   vi env
   . ./env
   ```

   * LOCATION: Location of the shared RP development environment (default:
     `eastus`).

1. Create the resource group and deploy the RP resources:

   ```bash
   . ./hack/devtools/deploy-shared-env.sh
   # Create the RG
   create_infra_rg
   # Deploy the predeployment ARM template
   deploy_rp_dev_predeploy
   # Deploy the infrastructure resources such as Cosmos, KV, Vnet...
   deploy_rp_dev
   # Deploy the proxy and VPN
   deploy_env_dev
   ```

   If you encounter a "VirtualNetworkGatewayCannotUseStandardPublicIP" error
   when running the `deploy_env_dev` command, you have to override two
   additional parameters.  Run this command instead:

   ```bash
   deploy_env_dev_override
   ```

   If you encounter a "SkuCannotBeChangedOnUpdate" error
   when running the `deploy_env_dev_override` command, delete the `-pip` resource
   and re-run.

1. Load the keys/certificates into the key vault:

   ```bash
   import_certs_secrets
   ```

   Note: in production, three additional keys/certificates (rp-mdm, rp-mdsd, and
   cluster-mdsd) are also required in the $KEYVAULT_PREFIX-svc key vault.  These
   are client certificates for RP metric and log forwarding (respectively) to
   Geneva.

   If you need them in development:

   ! DO NOT RUN THIS IN E2E, INT! !

   ```bash
   az keyvault certificate import \
        --vault-name "$KEYVAULT_PREFIX-svc" \
        --name rp-mdm \
        --file secrets/rp-metrics-int.pem
   az keyvault certificate import \
        --vault-name "$KEYVAULT_PREFIX-svc" \
        --name rp-mdsd \
        --file secrets/rp-logging-int.pem
   az keyvault certificate import \
        --vault-name "$KEYVAULT_PREFIX-svc" \
        --name cluster-mdsd \
        --file secrets/cluster-logging-int.pem
   ```

   Note: in development, if you don't have valid certs for these, you can just
   upload `localhost.pem` as a placeholder for each of these. This will avoid an
   error stemming from them not existing, but it will result in logging pods
   crash looping in any clusters you make.

1. In pre-production (int, e2e) certain certificates are provisioned via keyvault
integration. These should be rotated and generated in the keyvault itself:

```
Vault Name: "$KEYVAULT_PREFIX-svc"
Certificate: rp-firstparty
Development value: secrets/firstparty.pem

Vault Name: "$KEYVAULT_PREFIX-svc"
Certificate: cluster-mdsd
Development value: secrets/cluster-logging-int.pem
```

1. Create nameserver records in the parent DNS zone:

   ```bash
   update_parent_domain_dns_zone
   ```

1. Store the VPN client configuration:

   ```bash
   vpn_configuration
   ```
