<?php

namespace StormGeo\AdvisorCore\Payloads;

class GeometryPayload
{
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
   * @var string
   */
  public $geometry;

  /**
   * @var array
   */
  public function getQueryParams()
  {
    return [
      'startDate' => $this->startDate,
      'endDate' => $this->endDate,
      'radius' => $this->radius
    ];
  }

  /**
   * @var array
   */
  public function getBody()
  {
    return [
      'geometry' => $this->geometry
    ];
  }
}
