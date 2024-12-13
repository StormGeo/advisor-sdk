<?php

namespace StormGeo\AdvisorCore\Payloads;

class ClimatologyPayload
{
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
   * @var array
   */
  public function getQueryParams()
  {
    return (array) $this;
  }
}
