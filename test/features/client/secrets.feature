Feature: Secrets
  Secrets allows the system to write secrets.

  Scenario: Write existing secrets
    When I write secrets
    Then I should have secrets

  Scenario: Write missing secrets
    When I try to write missing secrets
    Then I should not have secrets
    And I should see a log entry of "secrets not found" in the file "reports/secrets.log"
