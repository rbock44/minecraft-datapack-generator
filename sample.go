package main

func newSample() *datapack {
	dp := &datapack{
		Name:         "simple-home",
		FunctionRoot: "true_home",
		Version:      1,
		Description:  "A set/go home function",
		TargetPath:   "generated",
	}

	namespace := "sub/home/go"
	dp.addFunction("error", namespace, `
tellraw @s [{"text":"","color":"green"}, {"text":"You do not have a home set.","color":"red"}, "\n", "Use ", {"text":"/trigger set_home","color":"gold"}, ", to set your home position."]
`)

	dp.addFunction("home", namespace, `
tag @s add home_player
summon minecraft:armor_stand ~ ~ ~ {Tags:["home_pos"],Invisible:1b,Marker:1b}
execute as @e[tag=home_pos,limit=1] run function true_home:sub/home/go/move
execute at @s if score @s home_d matches 0 in minecraft:overworld run tp @s ~0.5 ~ ~0.5
execute at @s if score @s home_d matches 1 in minecraft:the_end run tp @s ~0.5 ~ ~0.5
execute at @s if score @s home_d matches -1 in minecraft:the_nether run tp @s ~0.5 ~ ~0.5
tellraw @s ["", {"text":"You have arrived at your home.","color":"green"}]
scoreboard players reset @s go_home
tag @s remove home_player
`)

	dp.addFunction("move", namespace, `
execute store result entity @s Pos[0] double 1 run scoreboard players add @p[tag=home_player] home_x 0
execute store result entity @s Pos[1] double 1 run scoreboard players add @p[tag=home_player] home_y 0
execute store result entity @s Pos[2] double 1 run scoreboard players add @p[tag=home_player] home_z 0
execute at @s run tp @p[tag=home_player] ~ ~ ~
kill @s`)

	namespace = "sub/triggers"

	dp.addFunction("go_home", namespace, `
execute unless entity @s[tag=has_home] run function true_home:sub/home/go/error
execute if entity @s[tag=has_home] run function true_home:sub/home/go/home
scoreboard players reset @s go_home
`)

	dp.addFunction("set_home", namespace, `
execute store result score @s home_x run data get entity @s Pos[0]
execute store result score @s home_y run data get entity @s Pos[1]
execute store result score @s home_z run data get entity @s Pos[2]
execute store result score @s home_d run data get entity @s Dimension
scoreboard players reset @s set_home
tag @s add has_home
tellraw @s ["", {"text":"Your home has been set at","color":"green"}, ": ", {"score":{"name":"@s","objective":"home_x"},"color":"gold"}, ", ", {"score":{"name":"@s","objective":"home_y"},"color":"gold"}, ", ", {"score":{"name":"@s","objective":"home_z"},"color":"gold"}, "; ", {"score":{"name":"@s","objective":"home_d"},"color":"gold"}]
`)

	namespace = "sub/verification"

	dp.addFunction("validation", namespace, `
execute as @a[tag=has_home] unless score @s home_d matches ..0 unless score @s home_d matches 0.. run tag @s remove has_home
execute as @a[tag=has_home] unless score @s home_x matches ..0 unless score @s home_x matches 0.. run tag @s remove has_home
execute as @a[tag=has_home] unless score @s home_y matches ..0 unless score @s home_y matches 0.. run tag @s remove has_home
execute as @a[tag=has_home] unless score @s home_z matches ..0 unless score @s home_z matches 0.. run tag @s remove has_home
`)

	namespace = ""

	dp.addFunction("loop", namespace, `
function true_home:sub/verification/validation
function true_home:sub/triggers
`)

	dp.addFunction("setup", namespace, `
scoreboard objectives add set_home trigger
scoreboard objectives add go_home trigger
scoreboard objectives add home_d dummy
scoreboard objectives add home_x dummy
scoreboard objectives add home_y dummy
scoreboard objectives add home_z dummy
`)

	dp.addTag("load", `
{
  "values": [
    "true_home:setup"
  ]
}
`)

	dp.addTag("tick", `
{
  "values": [
    "true_home:loop"
  ]
}
`)

	return dp
}
