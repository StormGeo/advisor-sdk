<?php

namespace StormGeo\AdvisorCore\Payloads;

class RadiusPayload {
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
	public $startDate;

  /**
   * @var string
   */
  public $endDate;

  /**
   * @var int
   */
  public $radius;

  /**
   * @var array
   */
  public function getQueryParams()
  {
    return (array) $this;
  }
}
