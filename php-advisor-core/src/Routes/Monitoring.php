<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
class Monitoring extends BaseRouter
{
  /**
   * GET /v1/monitoring/alerts
   * 
   * @return  array
   */
  public function getAlerts()
  {
    return parent::makeRequest('GET', '/v1/monitoring/alerts' . $this->formatQueryParams());
  }
}
