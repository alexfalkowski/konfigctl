Feature: Secrets
  Secrets allows the system to write secrets.

  Scenario Outline: Write existing secrets
    When I write "<name>" secrets
    Then I should have secrets

    Examples: With different configs
      | name |
      | grpc |
      | http |

  Scenario Outline: Write missing secrets
    When I write "<name>" secrets
    Then I should not have secrets
    And I should see a log entry of "secrets not found" in the file "reports/secrets.log"

    Examples: With different configs
      | name         |
      | invalid_grpc |
      | invalid_http |
