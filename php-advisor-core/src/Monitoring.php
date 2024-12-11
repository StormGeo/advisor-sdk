<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
class Monitoring extends BaseRouter
{
  /**
   * @return  array
   */
  public function getAlerts()
  {
    return parent::makeRequest('GET', '/v1/monitoring/alerts');
  }
}
