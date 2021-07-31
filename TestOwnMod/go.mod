module github.com/ZeeeUs/GoTech2/TestOwnMod

go 1.16

replace github.com/ZeeeUs/GoTech2/OwnMod => ../OwnMod

require (
	github.com/ZeeeUs/GoTech2/OwnMod v0.0.0-00010101000000-000000000000
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/tools v0.1.5 // indirect
)
