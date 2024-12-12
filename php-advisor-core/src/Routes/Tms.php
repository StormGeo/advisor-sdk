<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
class Tms extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function get($payload)
  {
    return parent::makeRequest('GET', '/v1/tms');
  }
}
