<?php

namespace StormGeo\AdvisorCore\Routes;

/**
 * @package StormGeo\AdvisorCore
 */
class Observed extends BaseRouter
{
  /**
   * @param   array $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/daily', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getHourly($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/hourly', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getPeriod($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/period', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getLightning($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/lightning', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postLightning($payload)
  {
    return parent::makeRequest('POST', '/v1/observed/lightning', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getFireFocus($payload)
  {
    return parent::makeRequest('GET', '/v1/observed/fire-focus', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function postFireFocus($payload)
  {
    return parent::makeRequest('POST', '/v1/observed/fire-focus', $payload);
  }

  /**
   * @param   array $payload
   * @return  array
   */
  public function getStationData($payload)
  {
    return parent::makeRequest('GET', '/v1/station', $payload);
  }
}
