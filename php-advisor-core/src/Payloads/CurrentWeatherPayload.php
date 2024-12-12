<?php

namespace StormGeo\AdvisorCore\Payloads;

class CurrentWeatherPayload {
  /**
   * @var int
   */
	public $localeId;

  /**
   * @var string
   */
	public $latitude;

  /**
   * @var string
   */
	public $longitude;

  /**
   * @var string
   */
  public $stationId;

  /**
   * @var array<string>
   */
	public $variables;

  /**
   * @var int
   */
	public $timezone;

  /**
   * @var array
   */
  public function getQueryParams()
  {
    return (array) $this;
  }
}