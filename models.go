package uddf

type UDDF struct {
	Version             string               `xml:"version,attr"`
	Business            *Business            `xml:"business,omitempty"`
	DecoModel           *DecoModel           `xml:"decomodel,omitempty"`
	DiveComputerControl *DiveComputerControl `xml:"divecomputercontrol,omitempty"`
	Diver               Diver                `xml:"diver"`
	DiveSite            *DiveSite            `xml:"divesite,omitempty"`
	DiveTrip            *DiveTrip            `xml:"divetrip,omitempty"`
	GasDefinitions      *GasDefinitions      `xml:"gasdefinitions,omitempty" validate:"omitempty"`
	Generator           *Generator           `xml:"generator,omitempty"`
	Maker               *Maker               `xml:"maker,omitempty"`
	MediaData           *MediaData           `xml:"mediadata,omitempty"`
	ProfileData         ProfileData          `xml:"profiledata" validate:"required"`
	TableGeneration     TableGeneration      `xml:"tablegeneration"`
}

type DecoModel struct {
	Buehlmann []Buehlmann `xml:"buehlmann"` // parameter set for Bühlmann's decompression model
	RGBM      []RGBM      `xml:"rbgm"`      // parameter set for a Reduced Gradient Bubble Model (RGBM)
	VPM       []VPM       `xml:"vpm"`       // parameter set for a Varying Permeability Model (VPM)
}

type VPM struct {
	ID           string   `xml:"id,attr"`
	Conservatism *float64 `xml:"conservatism,omitempty"` // value of the respective Varying Permeability Model (VPM) parameter is set as a percentage. 42% == 0.42
	Gamma        *float64 `xml:"gamma,omitempty"`        // is the skin tension of bubble nuclei. units used for gamma are kg/s2
	GC           *float64 `xml:"gc,omitempty"`           // is the nuclear crushing tension. units used for gc are kg/s2
	Lambda       *float64 `xml:"lambda,omitempty"`       // denotes a summary of several magnitudes. units used for lambda are kg/m/s (7180 fsw*min = 367431.06061 kg/m/s)
	R0           *float64 `xml:"r0,omitempty"`           // minimum bubble radius excitable into growth. units used for r0 are metre
	Tissues      []Tissue `xml:"tissue"`                 // At least one <tissue/> element must appear inside the respective parent element
}

type RGBM struct {
	ID      string   `xml:"id,attr"`
	Tissues []Tissue `xml:"tissue"` // At least one <tissue/> element must appear inside the respective parent element
}

type Buehlmann struct {
	ID                 string   `xml:"id,attr"`
	GradientFactorHigh *float64 `xml:"gradientfactorhigh,omitempty"` // "Gradient Factor High" (GF High), given as a real number 0.0 <= GF Low <= GF High <= 1.0.
	GradientFactorLow  *float64 `xml:"gradientfactorlow,omitempty"`  // "Gradient Factor Low" (GF Low), given as a real number 0.0 <= GF Low <= GF High <= 1.0.
	Tissues            []Tissue `xml:"tissue"`                       //  At least one <tissue/> element must appear inside the respective parent element
}

type Tissue struct {
	Gas      string  `xml:"gas,attr"`      // h2 | he | n2
	HalfLife float64 `xml:"halflife,attr"` // halflife of the tissue, given in seconds as a real number
	Number   int     `xml:"number,attr"`   // number of tissue, given as an integer
	A        float64 `xml:"a,attr"`        // A value for the Buehlmann algorithm
	B        float64 `xml:"b,attr"`        // B value for the Buehlmann algorithm
}

type GasDefinitions struct {
	Mixes []Mix `xml:"mix,omitempty" validate:"dive"`
}

type Mix struct {
	ID                    string   `xml:"id,attr"`
	AliasName             *string  `xml:"aliasname"`
	Ar                    *float64 `xml:"ar,omitempty" validate:"omitempty,min=0,max=1"` // argon fraction of a (breathing) gas, given as a real number less or equal 1.0 in percent
	EquivalentAirDepth    *float64 `xml:"equivalentairdepth,omitempty"`                  // equivalent air depth of a gas, given as a real number in meters
	H2                    *float64 `xml:"h2,omitempty" validate:"omitempty,min=0,max=1"` // hydrogen fraction of a (breathing) gas, given as a real number less or equal 1.0 in percent
	He                    *float64 `xml:"he,omitempty" validate:"omitempty,min=0,max=1"` // helium fraction of a (breathing) gas, given as a real number less or equal 1.0 in percent
	MaximumOperationDepth *float64 `xml:"maximumoperationdepth,omitempty"`               // maximum operation depth of a gas, given as a real number in meters
	MaximumPo2            *float64 `xml:"maximumpo2,omitempty"`                          // threshold for the oxygen partial pressure, when the oxygen fraction of this breathing gas starts to be poisonous
	N2                    *float64 `xml:"n2,omitempty" validate:"omitempty,min=0,max=1"` // nitrogen fraction of a (breathing) gas, given as a real number less or equal 1.0 in percent
	Name                  string   `xml:"name" validate:"required"`
	O2                    *float64 `xml:"o2,omitempty" validate:"omitempty,min=0,max=1"` // oxygen fraction of a (breathing) gas, given as a real number less or equal 1.0 in percent
	PricePerLitre         *Price   `xml:"priceperlitre,omitempty"`
}

type ProfileData struct {
	RepetitionGroup []RepetitionGroup `xml:"repetitiongroup" validate:"dive"`
}

type RepetitionGroup struct {
	ID    *string `xml:"id,attr"`
	Dives []Dive  `xml:"dive" validate:"dive"`
}

type Dive struct {
	ID                    string                `xml:"id,attr"`
	InformationAfterDive  InformationAfterDive  `xml:"informationafterdive"`
	InformationBeforeDive InformationBeforeDive `xml:"informationbeforedive"`
	ApplicationData       *ApplicationData      `xml:"applicationdata,omitempty"`
	Samples               *Samples              `xml:"samples,omitempty"`
	TankData              []TankData            `xml:"tankdata"`
}

type InformationBeforeDive struct {
	AirTemperature            *float64                   `xml:"airtemperature,omitempty"`
	AlcoholBeforeDive         *AlcoholBeforeDive         `xml:"alcoholbeforedive,omitempty"`
	Altitude                  *float64                   `xml:"altitude,omitempty"`
	Apparatus                 *string                    `xml:"apparatus,omitempty"` // Allowed keywords are: open-scuba, rebreather, surface-supplied, chamber, experimental, other.
	DateTime                  Time                       `xml:"datetime"`
	DiveNumber                *int                       `xml:"diveNumber,omitempty"`
	DiveNumberOfDay           *int                       `xml:"diveNumberOfDay,omitempty"`
	InternalDiveNumber        *int                       `xml:"internalDiveNumber,omitempty"`
	Links                     []Link                     `xml:"link"`
	MedicationBeforeDive      *MedicationBeforeDive      `xml:"medicationbeforedive,omitempty"`
	NoSuit                    bool                       `xml:"nosuit"`
	PlannedProfile            *PlannedProfile            `xml:"plannedprofile,omitempty"`
	Platform                  *string                    `xml:"platform,omitempty"` // Allowed keywords are: beach-shore, pier, small-boat, charter-boat, live-aboard, barge, landside, hyperbaric-facility, other.
	Price                     *Price                     `xml:"price,omitempty"`
	Purpose                   *string                    `xml:"purpose,omitempty"`               // Allowed keywords are: sightseeing, learning, research, photography-videography, spearfishing, proficiency, work, other.
	StateOfRestBeforeDive     *string                    `xml:"stateofrestbeforedive,omitempty"` // Allowed keywords are: not-specified, rested, tired, exhausted.
	SurfaceIntervalBeforeDive *SurfaceIntervalBeforeDive `xml:"surfaceintervalbeforedive,omitempty"`
	SurfacePressure           *float64                   `xml:"surfacepressure,omitempty"`
	TripMembership            *string                    `xml:"tripmembership,omitempty"`
}

type MedicationBeforeDive struct {
	Medicines []Medicine `xml:"medicine,omitempty"`
}

type Medicine struct {
	AliasName          *string  `xml:"aliasname"`
	Name               string   `xml:"name"`
	Notes              *Notes   `xml:"notes,omitempty"`
	PeriodicallyTaken  *string  `xml:"periodicallytaken,omitempty"` // yes or no
	TimespanBeforeDive *float64 `xml:"timespanbeforedive,omitempty"`
}

type PlannedProfile struct {
	StartDiveMode string     `xml:"startdivemode,attr"`
	StartMix      string     `xml:"startmix,attr"`
	Waypoints     []Waypoint `xml:"waypoint"`
}

type AlcoholBeforeDive struct {
	Drinks []Drink `xml:"drink"`
}

type Drink struct {
	AliasName          *string  `xml:"aliasname"`
	Name               string   `xml:"name"`
	Notes              *Notes   `xml:"notes,omitempty"`
	PeriodicallyTaken  *string  `xml:"periodicallytaken,omitempty"` // yes or no
	TimespanBeforeDive *float64 `xml:"timespanbeforedive,omitempty"`
}

type InformationAfterDive struct {
	AnySymptoms              *AnySymptoms              `xml:"anysymptoms,omitempty"`
	AverageDepth             *float64                  `xml:"averagedepth,omitempty"`
	Current                  *string                   `xml:"current,omitempty"`              // Allowed keywords are: no-current, very-mild-current, mild-current, moderate-current, hard-current, very-hard-current.
	DesaturationTime         *float64                  `xml:"desaturationtime,omitempty"`     //
	DiveDuration             float64                   `xml:"diveduration"`                   //
	DivePlan                 *string                   `xml:"diveplan,omitempty"`             // Allowed keywords are: none, table, dive-computer, another-diver.
	DiveTable                *string                   `xml:"dive-table,omitempty"`           // Allowed keywords are: PADI, NAUI, BSAC, Buehlmann, DCIEM, US-Navy, CSMD, COMEX, other.
	EquipmentMalfunction     *string                   `xml:"equipmentmalfunction,omitempty"` // Allowed keywords are: none, face-mask, fins, weight-belt, buoyancy-control-device, thermal-protection (suit), dive-computer, depth-gauge, pressure-gauge, breathing-apparatus, deco-reel, other.
	EquipmentUsed            *EquipmentUsed            `xml:"equipmentused,omitempty"`
	GlobalAlarmsGiven        *GlobalAlarmsGiven        `xml:"globalalarmsgiven,omitempty"`
	GreatestDepth            float64                   `xml:"greatestdepth"`
	HighestPo2               *float64                  `xml:"highestpo2,omitempty"`
	LowestTemperature        *float64                  `xml:"lowesttemperature,omitempty"`
	NoFlightTime             *float64                  `xml:"noflighttime,omitempty"`
	Notes                    *Notes                    `xml:"notes,omitempty"`
	Observations             *Observations             `xml:"observations,omitempty"`
	PressureDrop             *float64                  `xml:"pressuredrop,omitempty"`
	Problems                 []string                  `xml:"problems,omitempty" validate:"dive,omitempty,oneof=none equalisation vertigo out-of-air buoyancy shared-air rapid-ascent sea-sickness other"`
	Program                  *string                   `xml:"program,omitempty" validate:"omitempty,oneof=recreation training scientific medical commercial military competitive other"`
	Ratings                  []Rating                  `xml:"rating,omitempty"`
	SurfaceIntervalAfterDive *SurfaceIntervalAfterDive `xml:"surfaceintervalafterdive,omitempty"`
	ThermalComfort           *string                   `xml:"thermalcomfort,omitempty"` // Allowed keywords are: not-indicated, comfortable, cold, very-cold, hot.
	Visibility               *FlexibleFloat            `xml:"visibility,omitempty"`
	Workload                 *string                   `xml:"workload,omitempty"` // Allowed keywords are: not-specified, resting, light, moderate, severe, exhausting.
}

type Observations struct {
	Fauna *Fauna `xml:"fauna,omitempty"`
	Flora *Flora `xml:"flora,omitempty"`
	Notes *Notes `xml:"notes,omitempty"`
}

type GlobalAlarmsGiven struct {
	GlobalAlarms []string `xml:"globalalarm,omitempty"` // ascent-warning-too-long, sos-mode, work-too-hard
}

type EquipmentUsed struct {
	LeadQuantity *float64 `xml:"leadquantity,omitempty"`
	Links        []Link   `xml:"link,omitempty"`
}

type AnySymptoms struct {
	Notes *Notes `xml:"notes,omitempty"`
}

type Samples struct {
	Waypoints []Waypoint `xml:"waypoint,omitempty"`
}

type TankData struct {
	ID                         string   `xml:"id,attr"`
	BreathingConsumptionVolume *float64 `xml:"breathingconsumptionvolume,omitempty"`
	Links                      []Link   `xml:"link,omitempty"`
	TankPressureBegin          float64  `xml:"tankpressurebegin"`
	TankPressureEnd            float64  `xml:"tankpressureend"`
	TankVolume                 *float64 `xml:"tankvolume,omitempty"` // Volume of the tank used in cubicmetres [m^3]
}

type TableGeneration struct {
	CalculateBottomTimeTable *CalculateBottomTimeTable `xml:"calculatebottomtimetable,omitempty"`
	CalculateProfile         *CalculateProfile         `xml:"calculateprofile,omitempty"`
	CalculateTable           *CalculateTable           `xml:"calculatetable,omitempty"`
}

type CalculateTable struct {
	Tables []Table `xml:"table,omitempty"`
}

type Table struct {
	Profile
	TableScope TableScope `xml:"tablescope"`
}

type TableScope struct {
	Altitude            *float64 `xml:"altitude,omitempty"`
	BottomTimeMaximum   *float64 `xml:"bottomtimemaximum,omitempty"`
	BottomTimeMinimum   *float64 `xml:"bottomtimeminimum,omitempty"`
	BottomTimeStepBegin *float64 `xml:"bottomtimestepbegin,omitempty"`
	BottomTimeStepEnd   *float64 `xml:"bottomtimestepend,omitempty"`
	DiveDepthBegin      *float64 `xml:"divedepthbegin,omitempty"`
	DiveDepthEnd        *float64 `xml:"divedepthend,omitempty"`
	DiveDepthStep       *float64 `xml:"divedepthstep,omitempty"`
}

type CalculateProfile struct {
	Profiles []Profile `xml:"profile"`
}

type Profile struct {
	ApplicationData           *ApplicationData           `xml:"applicationdata,omitempty"`
	DecoModel                 *DecoModel                 `xml:"decomodel,omitempty"`
	DeepStopTime              *float64                   `xml:"deepstoptime,omitempty"`
	Density                   *float64                   `xml:"density,omitempty"`
	InputProfile              *InputProfile              `xml:"inputprofile,omitempty"`
	Links                     []Link                     `xml:"link,omitempty"`
	MaximumAscendingRate      *float64                   `xml:"maximumascendingrate,omitempty"`
	MixChange                 MixChange                  `xml:"mixchange"`
	Output                    *Output                    `xml:"output,omitempty"`
	SurfaceIntervalAfterDive  *SurfaceIntervalAfterDive  `xml:"surfaceintervalafterdive,omitempty"`
	SurfaceIntervalBeforeDive *SurfaceIntervalBeforeDive `xml:"surfaceintervalbeforedive,omitempty"`
	Title                     *string                    `xml:"title,omitempty"`
}

type SurfaceIntervalBeforeDive struct {
	ExposureToAltitude *ExposureToAltitude `xml:"exposuretoaltitude,omitempty"`
	Infinity           *bool               `xml:"infinity,omitempty"`
	PassedTime         *float64            `xml:"passedtime,omitempty"`
	WayAltitudes       []WayAltitude       `xml:"wayaltitude"`
}

type SurfaceIntervalAfterDive struct {
	ExposureToAltitude *ExposureToAltitude `xml:"exposuretoaltitude,omitempty"`
	Infinity           *bool               `xml:"infinity,omitempty"`
	PassedTime         *float64            `xml:"passedtime,omitempty"`
	WayAltitudes       []WayAltitude       `xml:"wayaltitude,omitempty"`
}

type ExposureToAltitude struct {
	AltitudeOfExposure                    *float64 `xml:"altitudeofexposure,omitempty"`
	DateOfFlight                          *Date    `xml:"dateofflight,omitempty"`
	SurfaceIntervalBeforeAltitudeExposure *float64 `xml:"surfaceintervalbeforealtitudeexposure,omitempty"`
	TotalLengthOfExposure                 *float64 `xml:"totallengthofexposure,omitempty"`
	Transportation                        string   `xml:"transportation"` // Allowed keywords are: commercial-aircraft, unpressurized-aircraft, medevac-aircraft, ground-transportation, or helicopter.
}

type WayAltitude struct {
	WayTime float64 `xml:"waytime,attr"`
	Value   float64 `xml:",chardata"`
}

type MixChange struct {
	Ascent  Ascent  `xml:"ascent"`
	Descent Descent `xml:"descent"`
}

type Descent struct {
	Waypoints []Waypoint `xml:"waypoint"`
}

type Ascent struct {
	Waypoints []Waypoint `xml:"waypoint"`
}

type InputProfile struct {
	Links     []Link     `xml:"link"`
	Waypoints []Waypoint `xml:"waypoint"`
}

type Waypoint struct {
	Alarms                  []Alarm                  `xml:"alarm,omitempty"`
	BatteryChargeConditions []BatteryChargeCondition `xml:"batterychargecondition,omitempty"`
	CalculatedPo2           *float64                 `xml:"calculatedpo2,omitempty"`
	CNS                     *float64                 `xml:"cns,omitempty"`
	DecoStops               []Decostop               `xml:"decostop,omitempty"`
	Depth                   float64                  `xml:"depth"`
	DiveMode                *DiveMode                `xml:"divemode,omitempty"`
	DiveTime                float64                  `xml:"divetime"`
	GradientFactor          *GradientFactor          `xml:"gradientfactor,omitempty"`
	Heading                 *float64                 `xml:"heading,omitempty"`
	MeasuredPo2s            []MeasuredPo2            `xml:"measuredpo2"`
	NoDecoTime              *float64                 `xml:"nodecotime,omitempty"`
	OTU                     *float64                 `xml:"otu,omitempty"`
	RemainingBottomTime     *float64                 `xml:"remainingbottomtime,omitempty"`
	RemainingO2Time         *float64                 `xml:"remainingo2time,omitempty"`
	SetPo2s                 []SetPo2                 `xml:"setpo2"`
	SwitchMix               *SwitchMix               `xml:"switchmix,omitempty"`
	TankPressures           []TankPressure           `xml:"tankpressure,omitempty"`
	Temperature             float64                  `xml:"temperature"`
}

type TankPressure struct {
	Ref   *string `xml:"ref,attr,omitempty"`
	Value float64 `xml:",chardata"`
}

type SwitchMix struct {
	Ref string `xml:"ref,attr"`
}

type SetPo2 struct {
	SetBy string  `xml:"setby,attr"`
	Value float64 `xml:",chardata"`
}

type MeasuredPo2 struct {
	Ref   string  `xml:"ref,attr"`
	Value float64 `xml:",chardata"`
}

type GradientFactor struct {
	Tissue *int    `xml:"tissue,attr,omitempty"`
	Value  float64 `xml:",chardata"`
}

type DiveMode struct {
	Type string `xml:"type,attr"` // apnoe | closedcircuit | opencircuit | semiclosedcircuit
}

type Decostop struct {
	Kind      string  `xml:"kind,attr"`
	DecoDepth float64 `xml:"decodepth,attr"`
	Duration  float64 `xml:"duration,attr"`
}

type BatteryChargeCondition struct {
	DeviceRef string  `xml:"deviceref,attr"`
	TankRef   *string `xml:"tankref,attr,omitempty"`
	Value     float64 `xml:",chardata"`
}

type Alarm struct {
	Level   *float64 `xml:"level,attr,omitempty"`
	TankRef *string  `xml:"tankref,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type CalculateBottomTimeTable struct {
	BottomTimeTables []BottomTimeTable `xml:"bottomtimetable"`
}

type BottomTimeTable struct {
	ID                   string                `xml:"id,attr"`
	ApplicationData      *ApplicationData      `xml:"applicationdata,omitempty"`
	BottomTimeTableScope *BottomTimeTableScope `xml:"bottomtimetablescope,omitempty"`
	Links                []Link                `xml:"link"`
	Output               *Output               `xml:"output,omitempty"`
	Title                *string               `xml:"title,omitempty"`
}

type BottomTimeTableScope struct {
	BreathingConsumptionVolumeBegin *float64 `xml:"breathingconsumptionvolumebegin,omitempty"`
	BreathingConsumptionVolumeEnd   *float64 `xml:"breathingconsumptionvolumeend,omitempty"`
	BreathingConsumptionVolumeStep  *float64 `xml:"breathingconsumptionvolumestep,omitempty"`
	DiveDepthBegin                  *float64 `xml:"divedepthbegin,omitempty"`
	DiveDepthEnd                    *float64 `xml:"divedepthend,omitempty"`
	DiveDepthStep                   *float64 `xml:"divedepthstep,omitempty"`
	TankPressureBegin               *float64 `xml:"tankpressurebegin,omitempty"`
	TankPressureReserve             *float64 `xml:"tankpressurereserve,omitempty"`
	TankVolumeBegin                 *float64 `xml:"tankvolumebegin,omitempty"`
	TankVolumeEnd                   *float64 `xml:"tankvolumeend,omitempty"`
}

type Output struct {
	Lingo      *string `xml:"lingo,omitempty"`
	FileFormat *string `xml:"fileformat,omitempty"`
	FileName   *string `xml:"filename,omitempty"`
	Headline   *string `xml:"headline,omitempty"`
	Remark     *string `xml:"remark,omitempty"`
}

type ApplicationData struct {
	DecoTrainer *string   `xml:"decotrainer,omitempty"`
	Hargikas    *Hargikas `xml:"hargikas,omitempty"`
	// TODO: heinrichsweikamp
	// TODO: tausim
	// TODO: tautabu
}

type Hargikas struct {
	Ambient                      *float64 `xml:"ambient,omitempty"`                      // Ambient temperature when this dive starts (at 1.25m) (Celsius)
	Tissues                      []Tissue `xml:"tissue"`                                 // has eight <groups>
	ArterialMicroBubbleLevel     *int     `xml:"arterialmicrobubblelevel,omitempty"`     // the microbubble danger level in the arterial circulation (0 – 7)
	IntrapulmonaryRightLeftShunt *float64 `xml:"intrapulmonaryrightleftshunt,omitempty"` // Intrapulmonary right-left shunt: Micro bubbles in the venous circulation migrate to the lungs, where they are collected in the capillaries and obstruct the exchange of gas, and this effect is termed
	EstimatedSkinCoolLevel       *int     `xml:"estimatedskincoollevel,omitempty"`       //  skin cool level at dive start (0 – 7)
}

type MediaData struct {
	AudioFiles []Audio `xml:"audio,omitempty"`
	ImageFiles []Image `xml:"image,omitempty"`
	VideoFiles []Video `xml:"video,omitempty"`
}

type Video struct {
	ID         string  `xml:"id,attr"`
	ObjectName string  `xml:"objectname"`
	Title      *string `xml:"title,omitempty"`
}

type Image struct {
	ID         string     `xml:"id,attr"`
	Height     *int       `xml:"height,attr,omitempty"`
	Width      *int       `xml:"width,attr,omitempty"`
	Format     *string    `xml:"format,attr,omitempty"`
	ImageData  *ImageData `xml:"imagedata,omitempty"`
	ObjectName string     `xml:"objectname"`
	Title      *string    `xml:"title,omitempty"`
}

type ImageData struct {
	Aperture             *float64 `xml:"aperture,omitempty"`
	DateTime             *Time    `xml:"datetime,omitempty"`
	ExposureCompensation *float64 `xml:"exposurecompensation,omitempty"`
	FilmSpeed            *int     `xml:"filmspeed,omitempty"`
	FocalLength          *float64 `xml:"focallength,omitempty"`
	FocusingDistance     *float64 `xml:"focusingdistance,omitempty"`
	MeteringMethod       *string  `xml:"meteringmethod,omitempty"` // Allowed values are spot (spot metering), centerweighted (center-weighted metering), matrix (matrix metering).
	ShutterSpeed         *float64 `xml:"shutterspeed,omitempty"`
}

type Audio struct {
	ID         string  `xml:"id,attr"`
	ObjectName string  `xml:"objectname"`
	Title      *string `xml:"title"`
}

type Maker struct {
	Manufacturers []Manufacturer `xml:"manufacturer"`
}

type Generator struct {
	AliasName *string `xml:"aliasname,omitempty"`
	DateTime  *Time   `xml:"datetime,omitempty"`
	Links     []Link  `xml:"link"`
	Name      string  `xml:"name"`
	Type      *string `xml:"type,omitempty"` // Allowed keywords are: converter, divecomputer, and logbook.
	Version   *string `xml:"version,omitempty"`
}

type DiveTrip struct {
	Trips []Trip `xml:"trip"`
}

type Trip struct {
	ID        string     `xml:"id,attr"`
	AliasName *string    `xml:"aliasname,omitempty"`
	Name      string     `xml:"name"`
	Ratings   []Rating   `xml:"rating"`
	TripParts []TripPart `xml:"trippart"`
}

type TripPart struct {
	Type             *string           `xml:"type,attr,omitempty"`
	Accommodation    *Accommodation    `xml:"accommodation,omitempty"`
	DateOfTrip       *DateOfTrip       `xml:"dateoftrip,omitempty"`
	Geography        *Geography        `xml:"geography,omitempty"`
	Links            []Link            `xml:"link"`
	Name             string            `xml:"name"`
	Notes            *Notes            `xml:"notes,omitempty"`
	Operator         *Operator         `xml:"operator,omitempty"`
	PriceDivePackage *PriceDivePackage `xml:"pricedivepackage,omitempty"`
	PricePerDive     *Price            `xml:"priceperdive,omitempty"`
	RelatedDives     *RelatedDives     `xml:"relateddives,omitempty"`
	Vessel           *Vessel           `xml:"vessel,omitempty"`
}

type Vessel struct {
	Address       *Address       `xml:"address,omitempty"`
	AliasName     *string        `xml:"aliasname,omitempty"`
	Contact       *Contact       `xml:"contact,omitempty"`
	Marina        *string        `xml:"marina,omitempty"`
	Name          string         `xml:"name"`
	Notes         *Notes         `xml:"notes,omitempty"`
	Ratings       []Rating       `xml:"rating,omitempty"`
	ShipDimension *ShipDimension `xml:"shipdimension,omitempty"`
	ShipType      *string        `xml:"shiptype,omitempty"`
}

type RelatedDives struct {
	Links []Link `xml:"link"`
}

type Operator struct {
	AliasName *string  `xml:"aliasname,omitempty"`
	Address   *Address `xml:"address,omitempty"`
	Contact   *Contact `xml:"contact,omitempty"`
	Name      string   `xml:"name"`
	Notes     *Notes   `xml:"notes,omitempty"`
	Ratings   []Rating `xml:"rating,omitempty"`
}

type DateOfTrip struct {
	StartDate Time `xml:"startdate,attr"`
	EndDate   Time `xml:"enddate,attr"`
}

type Accommodation struct {
	Address   *Address `xml:"address,omitempty"`
	AliasName *string  `xml:"aliasname,omitempty"`
	Category  *string  `xml:"category,omitempty"`
	Contact   *Contact `xml:"contact,omitempty"`
	Name      string   `xml:"name"`
	Notes     *Notes   `xml:"notes,omitempty"`
	Ratings   []Rating `xml:"rating,omitempty"`
}

type DiveSite struct {
	DiveBases []DiveBase `xml:"divebase"`
	Sites     []Site     `xml:"site"`
}

type Site struct {
	ID          string     `xml:"id,attr"`
	AliasName   *string    `xml:"aliasname,omitempty"`
	Ecology     *Ecology   `xml:"ecology,omitempty"`
	Environment *string    `xml:"environment,omitempty"` // Allowed keywords are: unknown, ocean-sea, lake-quarry, river-spring, cave-cavern, pool, hyperbaric-chamber, under-ice, other.
	Geography   *Geography `xml:"geography,omitempty"`
	Links       []Link     `xml:"link,omitempty"`
	Name        string     `xml:"name"`
	Notes       *Notes     `xml:"notes,omitempty"`
	Ratings     []Rating   `xml:"rating,omitempty"`
	SideData    *SiteData  `xml:"sidedata,omitempty"`
}

type SiteData struct {
	AreaLength           *float64 `xml:"arealength,omitempty"`
	AreaWidth            *float64 `xml:"areawidth,omitempty"`
	AverageVisibility    *float64 `xml:"averagevisibility,omitempty"`
	Bottom               *string  `xml:"bottom,omitempty"`
	Cave                 *Cave    `xml:"cave,omitempty"`
	Density              *float64 `xml:"density,omitempty"`              //  Pure freshwater has a density of 1000.0 kg/m^3, whereas the mean density of sea water (salt water) is 1030.0 kg/m^3.
	Difficulty           *int     `xml:"difficulty,omitempty"`           // ranges from "1" (very easy to dive) through "10" (very difficult to dive).
	GlobalLightIntensity *string  `xml:"globallightintensity,omitempty"` // Allowed keywords are: undetermined, sunny, half-shadow, shadow, no-light (e.g. in a cave).
	Indoor               *Indoor  `xml:"indoor,omitempty"`
	MaximumDepth         *float64 `xml:"maximumdepth,omitempty"` // maximum depth in metres of the dive spot, not the greatest depth reached during a specific dive
	MaximumVisibility    *float64 `xml:"maximumvisibility,omitempty"`
	MinimumDepth         *float64 `xml:"minimumdepth,omitempty"`
	MinimumVisibility    *float64 `xml:"minimumvisibility,omitempty"`
	River                *River   `xml:"river,omitempty"`
	Shore                *Shore   `xml:"shore,omitempty"`
	Terrain              *string  `xml:"terrain,omitempty"`
	Wreck                *Wreck   `xml:"wreck,omitempty"`
}

type Wreck struct {
	AliasName     *string        `xml:"aliasname,omitempty"`
	Built         *Built         `xml:"built,omitempty"`
	Name          string         `xml:"name"`
	Nationality   *string        `xml:"nationality,omitempty"`
	ShipDimension *ShipDimension `xml:"shipdimension,omitempty"`
	ShipType      *string        `xml:"shiptype,omitempty"`
	Sunk          *Date          `xml:"sunk,omitempty"`
}

type ShipDimension struct {
	Beam         *float64 `xml:"beam,omitempty"`
	Displacement *float64 `xml:"displacement,omitempty"`
	Draught      *float64 `xml:"draught,omitempty"`
	Length       *float64 `xml:"length,omitempty"`
	Tonnage      *float64 `xml:"tonnage,omitempty"`
}

type Built struct {
	LaunchingDate Date    `xml:"launchingdate"`
	ShipYard      *string `xml:"shipyard,omitempty"`
}

type Shore struct {
	ID        string  `xml:"id,attr"`
	AliasName *string `xml:"aliasname,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
}

type River struct {
	ID        string  `xml:"id,attr"`
	AliasName *string `xml:"aliasname,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
}

type Lake struct {
	ID        string  `xml:"id,attr"`
	AliasName *string `xml:"aliasname,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
}

type Indoor struct {
	Address   *Address `xml:"address,omitempty"`
	AliasName *string  `xml:"aliasname,omitempty"`
	Contact   *Contact `xml:"contact,omitempty"`
	Name      string   `xml:"name"`
	Notes     *Notes   `xml:"notes,omitempty"`
}

type Cave struct {
	ID        string  `xml:"id,attr"`
	AliasName *string `xml:"aliasname,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
}

type Geography struct {
	Address   *Address `xml:"address,omitempty"`
	Altitude  *float64 `xml:"altitude,omitempty"`
	Latitude  *float64 `xml:"latitude,omitempty"`
	Location  string   `xml:"location"`
	Longitude *float64 `xml:"longitude,omitempty"`
	TimeZone  *float64 `xml:"timezone,omitempty"` // the difference to UTC in hours
}

type Ecology struct {
	Fauna *Fauna `xml:"fauna,omitempty"`
	Flora *Flora `xml:"flora,omitempty"`
}

type Flora struct {
	Chlorophyceae *WithSpecies `xml:"chlorophyceae,omitempty"`
	FloraVarious  *WithSpecies `xml:"floravarious,omitempty"`
	Notes         *Notes       `xml:"notes,omitempty"`
	Phaeophyceae  *WithSpecies `xml:"phaeophyceae,omitempty"`
	Rhodophyceae  *WithSpecies `xml:"rhodophyceae,omitempty"`
	Spermatophyta *WithSpecies `xml:"spermatophyta,omitempty"`
}

type Fauna struct {
	Invertebrata *Invertebrata `xml:"invertebrata,omitempty"`
	Notes        *Notes        `xml:"notes,omitempty"`
	Vertebrata   *Vertebrata   `xml:"vertebrata,omitempty"`
}

type Vertebrata struct {
	Amphibia          *WithSpecies `xml:"amphibia,omitempty"`
	Chondrichthyes    *WithSpecies `xml:"chondrichthyes,omitempty"`
	Mammalia          *WithSpecies `xml:"mammalia,omitempty"`
	Osteichthyes      *WithSpecies `xml:"osteichthyes,omitempty"`
	Reptilia          *WithSpecies `xml:"reptilia,omitempty"`
	VertebrataVarious *WithSpecies `xml:"vertebratavarious,omitempty"`
}

type Invertebrata struct {
	Ascidiacea          *WithSpecies `xml:"ascidiacea,omitempty"`
	Bryozoan            *WithSpecies `xml:"bryozoan,omitempty"`
	Cnidaria            *WithSpecies `xml:"cnidaria,omitempty"`
	Coelenterata        *WithSpecies `xml:"coelenterata,omitempty"`
	Crustacea           *WithSpecies `xml:"crustacea,omitempty"`
	Ctenophora          *WithSpecies `xml:"ctenophora,omitempty"`
	Echinodermata       *WithSpecies `xml:"echinodermata,omitempty"`
	InvertebrataVarious *WithSpecies `xml:"invertebratavarious,omitempty"`
	Mollusca            *WithSpecies `xml:"mollusca,omitempty"`
	Phoronidea          *WithSpecies `xml:"phoronidea,omitempty"`
	Plathelminthes      *WithSpecies `xml:"plathelminthes,omitempty"`
	Porifera            *WithSpecies `xml:"porifera,omitempty"`
}

type WithSpecies struct {
	Species []Species `xml:"species,omitempty"`
}

type Species struct {
	ID             string     `xml:"id,attr"`
	Abundance      *Abundance `xml:"abundance,omitempty"`
	Age            *int       `xml:"age,omitempty"`
	Dominance      *string    `xml:"dominance,omitempty"` // Allowed keywords are: undetermined, less-than-1/20, 1/20-up-to-1/4, 1/4-up-to-1/2, 1/2-up-to-3/4, greater-than-3/4, single-individual
	LifeStage      *string    `xml:"lifestage,omitempty"` // Keywords to be used are larva (larval stage), juvenile (young animal), or adult (adult animal)
	Notes          *Notes     `xml:"notes,omitempty"`
	ScientificName *string    `xml:"scientificname,omitempty"`
	Sex            *string    `xml:"sex,omitempty"` // keywords are: undetermined, male, female, hermaphrodite.
	Size           *float64   `xml:"size,omitempty"`
	TrivialName    *string    `xml:"trivialname,omitempty"`
}

type Abundance struct {
	Quality    *string `xml:"quality,attr,omitempty"`
	Occurrence *string `xml:"occurrence,attr,omitempty"`
	Value      int     `xml:",chardata"`
}

type DiveBase struct {
	ID               string            `xml:"id,attr"`
	Address          *Address          `xml:"address,omitempty"`
	AliasName        *string           `xml:"aliasname,omitempty"`
	Contact          *Contact          `xml:"contact,omitempty"`
	Guides           []Guide           `xml:"guide,omitempty"`
	Links            []Link            `xml:"link,omitempty"`
	Name             string            `xml:"name"`
	Notes            *Notes            `xml:"notes,omitempty"`
	PriceDivePackage *PriceDivePackage `xml:"pricedivepackage,omitempty"`
	PricePerDive     *Price            `xml:"priceperdive,omitempty"`
	Ratings          []Rating          `xml:"rating,omitempty"`
}

type Rating struct {
	DateTime    *Time `xml:"datetime,omitempty"`
	RatingValue int   `xml:"ratingvalue"` // The scale ranges from "1" (lowest quality) to "10" (highest quality).
}

type PriceDivePackage struct {
	Currency  string  `xml:"currency,attr"`
	NoOfDives int     `xml:"noofdives,attr"` // number of dives included in the package
	Value     float64 `xml:",chardata"`
}

type Guide struct {
	Links []Link `xml:"link"`
}

type Business struct {
	Shop *Shop `xml:"shop,omitempty"`
}

type Diver struct {
	Buddies []Buddy `xml:"buddy"`
	Owner   Owner   `xml:"owner"`
}

type BuddyOwnerShared struct {
	Id              string           `xml:"id,attr"`
	Address         *Address         `xml:"address,omitempty"`
	Contact         *Contact         `xml:"contact,omitempty"`
	DiveInsurances  *DiveInsurances  `xml:"diveinsurances,omitempty"`
	DivePermissions *DivePermissions `xml:"divepermissions,omitempty"`
	Equipment       *Equipment       `xml:"equipment,omitempty"`
	Medical         *Medical         `xml:"medical,omitempty"`
	Notes           *Notes           `xml:"notes,omitempty"`
	Personal        Personal         `xml:"personal"`
}

type Buddy struct {
	BuddyOwnerShared
	Certification *Certification `xml:"certification,omitempty"`
	Student       bool           `xml:"student"`
}

type Owner struct {
	BuddyOwnerShared
	Education *Education `xml:"education,omitempty"`
}

type Address struct {
	Street   *string `xml:"street,omitempty"`
	City     *string `xml:"city,omitempty"`
	Postcode *string `xml:"postcode,omitempty"`
	Country  string  `xml:"country"`
	Province *string `xml:"province,omitempty"`
}

type Personal struct {
	BirthDate     *Date          `xml:"birthdate,omitempty"`
	BirthName     *string        `xml:"birthname,omitempty"`
	BloodGroup    *string        `xml:"bloodgroup,omitempty"`
	FirstName     *string        `xml:"firstname,omitempty"`
	Height        *float64       `xml:"height,omitempty"`
	Honorific     *string        `xml:"honorific,omitempty"`
	LastName      *string        `xml:"lastname,omitempty"`
	Membership    *Membership    `xml:"membership,omitempty"`
	MiddleName    *string        `xml:"middlename,omitempty"`
	NumberOfDives *NumberOfDives `xml:"numberofdives,omitempty"`
	Sex           *string        `xml:"sex,omitempty"`
	Smoking       *string        `xml:"smoking,omitempty"`
	Weight        *float64       `xml:"weight,omitempty"` //  The element puts into brackets the weight (given in kilograms as a real number) of the owner of the UDDF file.
}

type Date struct {
	DateTime Time `xml:"datetime"`
}

type Membership struct {
	Organisation string  `xml:"organisation,attr"`
	MemberID     *string `xml:"memberid,attr,omitempty"`
}

type NumberOfDives struct {
	StartDate Date `xml:"startdate"`
	EndDate   Date `xml:"enddate"`
	Dives     int  `xml:"dives"`
}

type Education struct {
	Certifications []Certification `xml:"certification"`
}

type Certification struct {
	CertificateNumber *string     `xml:"certificatenumber,omitempty"`
	Instructor        *Instructor `xml:"instructor,omitempty"`
	IssueDate         *Date       `xml:"issuedate,omitempty"`
	Level             string      `xml:"level"`
	Link              *Link       `xml:"link,omitempty"`
	Organization      *string     `xml:"organization,omitempty"`
	Specialty         string      `xml:"specialty"`
	ValidDate         *Date       `xml:"validdate,omitempty"`
}

type Instructor struct {
	Address  *Address `xml:"address,omitempty"`
	Contact  *Contact `xml:"contact,omitempty"`
	Personal Personal `xml:"personal"`
}

type Link struct {
	Ref string `xml:"ref,attr"`
}

type Contact struct {
	Emails       []string `xml:"email"`
	Faxes        []string `xml:"fax"`
	Homepages    []string `xml:"homepage"`
	Languages    []string `xml:"language"`
	MobilePhones []string `xml:"mobilephone"`
	Phones       []string `xml:"phone"`
}

type Notes struct {
	Paras []string `xml:"para"`
	Links []Link   `xml:"link"`
}

type DiveInsurances struct {
	Insurances []Insurance `xml:"insurance"`
}

type Insurance struct {
	AliasName *string `xml:"aliasname"`
	IssueDate *Date   `xml:"issuedate,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
	ValidDate *Date   `xml:"validdate,omitempty"`
}

type DivePermissions struct {
	Permits []Permit `xml:"permit"`
}

type Permit struct {
	AliasName *string `xml:"aliasname"`
	IssueDate *Date   `xml:"issuedate,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
	Region    *string `xml:"region,omitempty"`
	ValidDate *Date   `xml:"validdate,omitempty"`
}

type Equipment struct {
	EquipmentContent
	Compressors            []EquipmentPart        `xml:"compressor"`
	EquipmentConfiguration EquipmentConfiguration `xml:"equipmentconfiguration"`
}

type EquipmentConfiguration struct {
	EquipmentContent
	AliasName *string `xml:"aliasname,omitempty"`
	Links     []Link  `xml:"link,omitempty"`
	Name      string  `xml:"name"`
	Notes     *Notes  `xml:"notes,omitempty"`
}

// To deduplicate equipment code
type EquipmentContent struct {
	Boots                  []EquipmentPart `xml:"boots"`
	BuoyancyControlDevices []EquipmentPart `xml:"buoyancycontroldevice"`
	Cameras                []EquipmentPart `xml:"camera"`
	Compasses              []EquipmentPart `xml:"compass"`
	DiveComputers          []EquipmentPart `xml:"divecomputer"`
	Fins                   []EquipmentPart `xml:"fins"`
	Gloves                 []EquipmentPart `xml:"gloves"`
	Knives                 []EquipmentPart `xml:"knife"`
	Leads                  []Lead          `xml:"lead"`
	Lights                 []EquipmentPart `xml:"light"`
	Masks                  []EquipmentPart `xml:"mask"`
	Rebreathers            []Rebreather    `xml:"rebreather"`
	Regulators             []EquipmentPart `xml:"regulator"`
	Scooters               []EquipmentPart `xml:"scooter"`
	Suits                  []Suit          `xml:"suit"`
	Tanks                  []Tank          `xml:"tank"`
	VariousPieces          []EquipmentPart `xml:"variouspieces"`
	VideoCameras           []EquipmentPart `xml:"videocamera"`
	Watches                []EquipmentPart `xml:"watch"`
}

type Tank struct {
	EquipmentPart
	TankMaterial *string  `xml:"tankmaterial,omitempty"` //  Indicates the material a tank is made of. Possible values are either aluminium, or carbon, or steel respectively.
	TankVolume   *float64 `xml:"tankvolume,omitempty"`   //  Volume of the tank used in cubicmetres [m^3] — not in litres, as UDDF uses SI units!
}

type Suit struct {
	EquipmentPart
	SuitType *string `xml:"suittype,omitempty"`
}

type Rebreather struct {
	EquipmentPart
	O2Sensors []EquipmentPart `xml:"o2sensor"`
}

type Lead struct {
	EquipmentPart
	LeadQuantity *int `xml:"leadquantity,omitempty"`
}

type EquipmentPart struct {
	Id              string        `xml:"id,attr"`
	AliasName       *string       `xml:"aliasname,omitempty"`
	Links           []Link        `xml:"link"`
	Manufacturer    *Manufacturer `xml:"manufacturer,omitempty"`
	Model           *string       `xml:"model,omitempty"`
	Name            string        `xml:"name"`
	NextServiceDate *Date         `xml:"nextservicedate,omitempty"`
	Notes           *Notes        `xml:"notes,omitempty"`
	Purchase        *Purchase     `xml:"purchase,omitempty"`
	SerialNumber    *string       `xml:"serialnumber,omitempty"`
	ServiceInterval *int          `xml:"serviceinterval,omitempty"`
}

type Purchase struct {
	DateTime *Time  `xml:"datetime,omitempty"`
	Link     *Link  `xml:"link,omitempty"`
	Price    *Price `xml:"price,omitempty"`
	Shop     *Shop  `xml:"shop,omitempty"`
}

type Shop struct {
	AliasName *string  `xml:"aliasname,omitempty"`
	Address   *Address `xml:"address,omitempty"`
	Contact   *Contact `xml:"contact,omitempty"`
	Name      string   `xml:"name"`
	Notes     *Notes   `xml:"notes,omitempty"`
}

type Price struct {
	Currency string  `xml:"currency,attr"`
	Value    float64 `xml:",chardata"`
}

type Manufacturer struct {
	Address   *Address `xml:"address,omitempty"`
	AliasName *string  `xml:"aliasname,omitempty"`
	Contact   *Contact `xml:"contact,omitempty"`
	Name      string   `xml:"name"`
}

type Camera struct {
	Body    EquipmentPart   `xml:"body"`
	Flashes []EquipmentPart `xml:"flash"`
	Housing *EquipmentPart  `xml:"housing,omitempty"`
	Lens    *EquipmentPart  `xml:"lens,omitempty"`
}

type Medical struct {
	Examination Examination `xml:"examination"`
}

type Examination struct {
	DateTime          *Time    `xml:"datetime,omitempty"`
	Doctor            *Doctor  `xml:"doctor,omitempty"`
	ExaminationResult *string  `xml:"examinationresult,omitempty"` // The only values allowed are passed, and failed respectively.
	Links             []Link   `xml:"link"`
	Notes             *Notes   `xml:"notes,omitempty"`
	TotalLungCapacity *float64 `xml:"totallungcapacity,omitempty"` // The value is given in m^3 as a real number. 6,4 litres == 0.0064
	VitalCapacity     *float64 `xml:"vitalcapacity,omitempty"`     // amount of air that can be forced out of the lungs after a maximal inspiration, given in m^3 as a real number. 4,5 litres == 0.0045
}

type Doctor struct {
	Id       string   `xml:"id,attr"`
	Address  *Address `xml:"address,omitempty"`
	Contact  *Contact `xml:"contact,omitempty"`
	Personal Personal `xml:"personal"`
}

type DiveComputerControl struct {
	DiveComputerDumps []DiveComputerDump `xml:"divecomputerdump"`
	GetDCData         *GetDCData         `xml:"getdcdata,omitempty"`
	SetDCData         *SetDCData         `xml:"setdcdata,omitempty"`
}

type SetDCData struct {
	SetDCAlarmTime          *Time                `xml:"setdcalarmtime,omitempty"`
	SetDCAltitude           *float64             `xml:"setdcaltitude,omitempty"`
	SetDCBuddyData          *SetDCBuddyData      `xml:"setdcbuddydata,omitempty"`
	SetDCDateTime           *Time                `xml:"setdcdatetime,omitempty"`
	SetDCDecoModel          *SetDCDecoModel      `xml:"setdcdecomodel,omitempty"`
	SetDCDiveDepthAlarm     *SetDCDiveDepthAlarm `xml:"setdcdepthalarm,omitempty"`
	SetDCDivePo2Alarm       *SetDCDivePo2Alarm   `xml:"setdcpo2alarm,omitempty"`
	SetDCDiveSiteData       []SetDCDiveSiteData  `xml:"setdcdivesitedata"`
	SetDCDiveTimeAlarm      *SetDCDiveTimeAlarm  `xml:"setdcitimealarm,omitempty"`
	SetDCEndNDTAlarm        *SetDCEndNDTAlarm    `xml:"setdcendndtalarm,omitempty"`
	SetDCGasDefinitionsData bool                 `xml:"setdcgasdefinitionsdata"`
	SetDCOwnerData          bool                 `xml:"setdcownerdata"`
	SetDCPassword           *string              `xml:"setdcpassword,omitempty"`
	SetDCGeneratorData      bool                 `xml:"setdcgeneratordata"`
}

type SetDCEndNDTAlarm struct {
	DCAlarm DCAlarm `xml:"dcalarm"`
}

type SetDCDiveTimeAlarm struct {
	DCAlarm  DCAlarm `xml:"dcalarm"`
	Timespan float64 `xml:"timespan"`
}

type SetDCDiveSiteData struct {
	DiveSite string `xml:"divesite,attr"`
}

type SetDCDivePo2Alarm struct {
	DCAlarm    DCAlarm  `xml:"dcalarm"`
	MaximumPo2 *float64 `xml:"maximumpo2,omitempty"`
}

type SetDCDiveDepthAlarm struct {
	DCAlarm      DCAlarm `xml:"dcalarm"`
	DCAlarmDepth float64 `xml:"dcalarmdepth"`
}

type DCAlarm struct {
	Acknowledge bool     `xml:"acknowledge"`
	AlarmType   int      `xml:"alarmtype"`
	Period      *float64 `xml:"period,omitempty"`
}

type SetDCDecoModel struct {
	AliasName       *string          `xml:"aliasname"`
	ApplicationData *ApplicationData `xml:"applicationdata,omitempty"`
	Name            string           `xml:"name"`
}

type SetDCBuddyData struct {
	Buddy string `xml:"buddy,attr"`
}

type GetDCData struct {
	GetDCAllData            bool `xml:"getdcalldata"`
	GetDCGeneratorData      bool `xml:"getdcgeneratordata"`
	GetDCOwnerData          bool `xml:"getdcownerdata"`
	GetDCBuddyData          bool `xml:"getdcbuddydata"`
	GetDCGasDefinitionsData bool `xml:"getdcgasdefinitionsdata"`
	GetDCDiveSiteData       bool `xml:"getdcdivesitedata"`
	GetDCDiveTripData       bool `xml:"getdcdivetripdata"`
	GetDCProfileData        bool `xml:"getdcprofiledata"`
}

type DiveComputerDump struct {
	DateTime Time   `xml:"datetime"`
	DCDump   string `xml:"dcdump"`
	Link     *Link  `xml:"link,omitempty"`
}
