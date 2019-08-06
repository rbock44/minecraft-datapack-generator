package main

import "fmt"

const (
	nsRoot         namespace = ""
	nsGoHome       namespace = "sub/home/go"
	nsTriggers     namespace = "sub/triggers"
	nsVerification namespace = "sub/verification"
)

const (
	funcNameLoop       functionName = "loop"
	funcNameSetup      functionName = "setup"
	funcNameValidation functionName = "validation"
	funcNameSetHome    functionName = "set_home"
	funcNameGoHome     functionName = "go_home"
	funcNameHome       functionName = "home"
	funcNameMove       functionName = "move"
	funcNameError      functionName = "error"
)

const (
	funcRefHome       functionRef = "true_home:sub/home/go/home"
	funcRefMove       functionRef = "true_home:sub/home/go/move"
	funcRefError      functionRef = "true_home:sub/home/go/error"
	funcRefValidation functionRef = "true_home:sub/verification/validation"
	funcRefTriggers   functionRef = "true_home:sub/triggers"
)

const (
	obSetHome objective = "set_home"
	obGoHome  objective = "go_home"
	obHomeD   objective = "home_d"
	obHomeX   objective = "home_x"
	obHomeY   objective = "home_y"
	obHomeZ   objective = "home_z"
)

var msgError = newMessageList(
	message{Color: "green"},
	message{Text: "You do not have a home set.", Color: "red"},
	message{Word: "\n"},
	message{Word: "Use "},
	message{Text: "/trigger set_home", Color: "gold"},
	message{Word: ", to set your home position."})

var msgSetHome = newMessageList(
	message{Text: "Your home has been set at", Color: "green"},
	message{Word: ""},
	message{Text: "", Color: "green"},
	message{Word: ": "},
	message{Word: `{"score":{"name":"@s","objective":"home_x"},"color":"gold"}`},
	message{Word: `{"score":{"name":"@s","objective":"home_y"},"color":"gold"}`},
	message{Word: ", "},
	message{Word: `{"score":{"name":"@s","objective":"home_z"},"color":"gold"}`},
	message{Word: `{"score":{"name":"@s","objective":"home_d"},"color":"gold"}`})

var msgArrivedAtHome = newMessageList(
	message{Word: ""},
	message{Text: "You have arrived at your home.", Color: "green"})

var triggerObjectives = []objective{obSetHome, obGoHome}
var dummyObjectives = []objective{obHomeD, obHomeX, obHomeY, obHomeZ}

func newSample() {
	dp = &datapack{
		Name:         "simple-home",
		FunctionRoot: "true_home",
		Version:      1,
		Description:  "A set/go home function",
		TargetPath:   "generated",
	}

	sampleError()
	sampleHome()
	sampleMove()
	sampleGoHome()
	sampleSetHome()
	sampleVerification()
	sampleLoop()
	sampleSetup()

	dp.addLoadTag("true_home:setup")
	dp.addTickTag("true_home:loop")

	dp.generate()
}

func sampleError() {
	fc := dp.newFunction(funcNameError, nsGoHome).Content
	fc.tellRaw(msgError)
}

func sampleHome() {
	fc := dp.newFunction(funcNameHome, nsGoHome).Content
	fc.addHomePlayer()
	fc.addLine(`summon minecraft:armor_stand ~ ~ ~ {Tags:["home_pos"],Invisible:1b,Marker:1b}`)
	fc.executeAs("@e[tag=home_pos,limit=1]", "run function true_home:sub/home/go/move")
	fc.executeAt("@s", "if score @s home_d matches 0 in minecraft:overworld run tp @s ~0.5 ~ ~0.5")
	fc.executeAt("@s", "if score @s home_d matches 1 in minecraft:the_end run tp @s ~0.5 ~ ~0.5")
	fc.executeAt("@s", "if score @s home_d matches -1 in minecraft:the_nether run tp @s ~0.5 ~ ~0.5")
	fc.tellRaw(msgArrivedAtHome)
	fc.scoreReset("@s", obGoHome)
	fc.removeHomePlayer()
}

func sampleMove() {
	fc := dp.newFunction(funcNameMove, nsGoHome).Content
	format := "result entity @s %s double 1 run scoreboard players add @p[tag=home_player] %s 0"
	fc.executeStore(fmt.Sprintf(format, "Pos[0]", obHomeX))
	fc.executeStore(fmt.Sprintf(format, "Pos[1]", obHomeY))
	fc.executeStore(fmt.Sprintf(format, "Pos[2]", obHomeZ))
	fc.executeAt("@s", "run tp @p[tag=home_player] ~ ~ ~")
	fc.addLine(`kill @s`)

}

func sampleGoHome() {
	fc := dp.newFunction(funcNameGoHome, nsTriggers).Content
	fc.executeUnless("entity @s[tag=has_home] run function " + string(funcRefError))
	fc.executeIf("entity @s[tag=has_home] run function " + string(funcRefHome))
	fc.scoreReset("@s", obGoHome)

}

func sampleSetHome() {
	fc := dp.newFunction(funcNameSetHome, nsTriggers).Content
	executeStore := "result score @s %s run data get entity @s %s"
	fc.executeStore(fmt.Sprintf(executeStore, obHomeX, "Pos[0]"))
	fc.executeStore(fmt.Sprintf(executeStore, obHomeY, "Pos[1]"))
	fc.executeStore(fmt.Sprintf(executeStore, obHomeZ, "Pos[2]"))
	fc.executeStore(fmt.Sprintf(executeStore, obHomeD, "Dimension"))
	fc.scoreReset("@s", obSetHome)
	fc.addLine(`tag @s add has_home`)
	fc.tellRaw(msgSetHome)
}

func sampleVerification() {
	fc := dp.newFunction(funcNameValidation, nsVerification).Content
	for _, ob := range []objective{obHomeD, obHomeX, obHomeY, obHomeZ} {
		fc.executeAs("@a[tag=has_home]", fmt.Sprintf("unless score @s %s matches ..0 unless score @s %s matches 0.. run tag @s remove has_home", ob, ob))
	}
}

func sampleLoop() {
	fc := dp.newFunction(funcNameLoop, nsRoot).Content
	fc.addFunctionRef(funcRefValidation)
	fc.addFunctionRef(funcRefTriggers)
}

func sampleSetup() {
	fc := dp.newFunction(funcNameSetup, nsRoot).Content
	for _, ob := range triggerObjectives {
		fc.addObjective(ob, criteriaTrigger, "")
	}

	for _, ob := range dummyObjectives {
		fc.addObjective(ob, criteriaDummy, "")
	}
}
