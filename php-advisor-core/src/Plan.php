<?php

namespace StormGeo\AdvisorCore;

/**
 * @package StormGeo\AdvisorCore
 */
class Plan extends BaseRouter
{
  /**
   * @return  array
   */
  public function getInfo()
  {
    return parent::makeRequest('GET', '/v1/plan/:token');
  }
}
