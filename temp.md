well, it works but i want to make sure i understand everything.
0. the main.go file in the rom_generator dir creates the ROM, which is a file that contains all the commands/instructions that the gameboy will execute.
1. main.go: we setup ebitengine's functions to render the gameboy's screen.
2. we create a new gameboy instance on line 66 `gb := NewGameBoy()`, which creates a variable Gameboy with the CPU's PC (pointer to the next instruction to be executed) at 0x0100.
3. we load the rom into the gameboy. i am also a bit lost here. what does 'load the rom' actually do? does it load the ROM instructions in a place in the MMU? (memory addresses 0x0000 all the way to 0x7FFF?)
4. this syntax is a bit confusing for me, who is new to go:
@main.go (67-70) 
explain please
5 game := &Game{gb: gb} we create the Game from ebitengine with the gameboy variable which it takes as an argument.
6. game starts running, so the Update() cycle starts, which runs 60 times per second if i am not mistaken. if i am correct, how does this 
@main.go (26-28) 
tie to ebitengine's Draw()?
7. every time it runs, it calls Draw() to draw the new frame to the screen. (right now it just create a black background)
8. it executes the Step function, which fetches the next instruction loaded from the ROM, decodes it and then executes it. it does so by reading the current instruction in the MMU, modifying the CPU accordingly and making any other change necessary and then jumps to the next instruction in the ROM by advancing the PC.
synopsis:
- correct me on my flow where i am mistaken
- explain this syntax: @main.go (67-69) 
- the screen is black because of this part of the files, right? @main.go (16-17)  because the ROM loads the value 0x03 in the FF47 register (load high A, 0x03)
- explain correlation between Draw(), Update() and the gameboy's fps