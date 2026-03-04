terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.62.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "pjs_tools" {
  name     = "pjs-tools-rg"
  location = "Australia East"
}

resource "azurerm_user_assigned_identity" "pjs_tools" {
  location            = azurerm_resource_group.pjs_tools.location
  name                = "pjs-tools-id"
  resource_group_name = azurerm_resource_group.pjs_tools.name
}

resource "azurerm_role_assignment" "pjs_tools" {
  scope                = azurerm_linux_function_app.pjs_tools.id
  role_definition_name = "Website Contributor"
  principal_id         = azurerm_user_assigned_identity.pjs_tools.principal_id
}


resource "azurerm_federated_identity_credential" "pjs_tools" {
  name      = "pjs_tools"
  audience  = ["api://AzureADTokenExchange"]
  issuer    = "https://token.actions.githubusercontent.com"
  parent_id = azurerm_user_assigned_identity.pjs_tools.id
  subject   = "repo:pjsmith404/pjs-tools:ref:refs/heads/master"
}

resource "azurerm_service_plan" "pjs_tools" {
  name                = "pjs_tools"
  resource_group_name = azurerm_resource_group.pjs_tools.name
  location            = azurerm_resource_group.pjs_tools.location
  os_type             = "Linux"
  sku_name            = "Y1"
}

resource "azurerm_storage_account" "pjs_tools" {
  name                     = "pjstoolsstorageaccount"
  resource_group_name      = azurerm_resource_group.pjs_tools.name
  location                 = azurerm_resource_group.pjs_tools.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_linux_function_app" "pjs_tools" {
  name                = "pjsTools"
  resource_group_name = azurerm_resource_group.pjs_tools.name
  location            = azurerm_resource_group.pjs_tools.location

  storage_account_name       = azurerm_storage_account.pjs_tools.name
  storage_account_access_key = azurerm_storage_account.pjs_tools.primary_access_key
  service_plan_id            = azurerm_service_plan.pjs_tools.id

  site_config {}
}
