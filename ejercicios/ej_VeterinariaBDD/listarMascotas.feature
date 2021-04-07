Feature: Listar todas las mascotas con sus datos

  Scenario:
    Given Existe un listado de mascotas con 4 registros
    When se solicita listar las mascotas
    Then el listado que aparece es de 4 mascotas

  Scenario:
    Given No hay un listado de mascotas
    When se solicita listar las mascotas
    Then el listado que aparece es 0 mascotas
    And genera error "No existen mascotas registradas"
