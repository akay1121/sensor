package entity

import "time"

// SensorDataProcessing reprsents the entity of it
type SensorDataProcessing struct {
	ID          string    // snsor ID
	TypeID      int64     // sensor type Id
	Description string    // description
	RuleID      int64     // rule ID
	Status      string    // sensor status
	LastUpdate  time.Time // last update time
	// i don't konw whether need i to make the rule ,if don't just delete it
	Threshold     float64 // 设定的阈值
	PreviousValue float64 // 之前的传感器数据值

}
