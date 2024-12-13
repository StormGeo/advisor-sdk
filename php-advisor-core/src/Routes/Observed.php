<?php

namespace StormGeo\AdvisorCore\Routes;

use StormGeo\AdvisorCore\Payloads\BasePayload;
use StormGeo\AdvisorCore\Payloads\GeometryPayload;
use StormGeo\AdvisorCore\Payloads\StationPayload;
use StormGeo\AdvisorCore\Payloads\RadiusPayload;

/**
 * @package StormGeo\AdvisorCore
 */
class Observed extends BaseRouter
{
  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getDaily($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/observed/daily' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getHourly($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/observed/hourly' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   BasePayload $payload
   * @return  array
   */
  public function getPeriod($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/observed/period' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   RadiusPayload $payload
   * @return  array
   */
  public function getLightning($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/observed/lightning' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   GeometryPayload $payload
   * @return  array
   */
  public function postLightning($payload)
  {
    return parent::makeRequest(
      'POST',
      '/v1/observed/lightning' . $this->formatQueryParams($payload->getQueryParams()),
      $payload->getBody()
    );
  }

  /**
   * @param   RadiusPayload $payload
   * @return  array
   */
  public function getFireFocus($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/observed/fire-focus' . $this->formatQueryParams($payload->getQueryParams())
    );
  }

  /**
   * @param   GeometryPayload $payload
   * @return  array
   */
  public function postFireFocus($payload)
  {
    return parent::makeRequest(
      'POST',
      '/v1/observed/fire-focus' . $this->formatQueryParams($payload->getQueryParams()),
      $payload->getBody()
    );
  }

  /**
   * @param   StationPayload $payload
   * @return  array
   */
  public function getStationData($payload)
  {
    return parent::makeRequest(
      'GET',
      '/v1/station' . $this->formatQueryParams($payload->getQueryParams())
    );
  }
}
