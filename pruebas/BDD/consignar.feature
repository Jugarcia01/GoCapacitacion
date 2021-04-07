
Feature: Consignar un monto a una cuenta

  Scenario:
    Given la cuenta tiene un monto de 10000
    When se consigna 2000
    Then la cuenta tiene un monto de 12000

  Scenario:
    Given la cuenta tiene un monto de 10000
    When se consigna -2000
    Then la cuenta tiene un monto de 10000
    And genera error "El valor no debe ser negativo"