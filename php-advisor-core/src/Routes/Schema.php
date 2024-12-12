<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
class Schema extends BaseRouter
{
  /**
   * @return  array
   */
  public function getDefinition()
  {
    return parent::makeRequest('GET', '/v1/schema/definition');
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postDefinition($payload)
  {
    return parent::makeRequest('POST', '/v1/schema/definition', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postParameters($payload)
  {
    return parent::makeRequest('POST', '/v1/schema/parameters', $payload);
  }
}
