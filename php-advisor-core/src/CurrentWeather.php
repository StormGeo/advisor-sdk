<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
class CurrentWeather extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function get($payload)
  {
    return parent::makeRequest('GET', '/v1/current-weather', $payload);
  }
}
