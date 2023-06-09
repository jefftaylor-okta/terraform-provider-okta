resource "okta_app_oauth" "test" {
  label                  = "testAcc_replace_with_uuid"
  status                 = "ACTIVE"
  type                   = "browser"
  grant_types            = ["authorization_code"]
  redirect_uris          = ["http://d.com/aaa"]
  response_types         = ["code"]
  hide_ios               = true
  hide_web               = true
  auto_submit_toolbar    = false
  refresh_token_rotation = "ROTATE"
  refresh_token_leeway   = 30
}
