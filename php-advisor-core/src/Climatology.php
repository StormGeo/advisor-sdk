<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
class Climatology extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/climatology/daily', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getMonthly($payload)
  {
    return parent::makeRequest('GET', '/v1/climatology/monthly', $payload);
  }
}
