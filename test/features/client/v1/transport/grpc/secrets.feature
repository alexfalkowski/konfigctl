Feature: Secrets

  Secrets allows the system to write secrets.

  Scenario: Write existing secrets
    When I write secrets
    Then I should have secrets
