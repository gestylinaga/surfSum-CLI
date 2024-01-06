package surfwords

import (
  "math/rand"
)

func Words() string{
  wordList := []string {
    "A-frame",
    "Aggro",
    "Air",
    "Alaia",
    "Aloha",
    "Amped",
    "Ankle slappers",
    "Backdoor",
    "Backside",
    "Backwash",
    "Bad vibes",
    "Bail",
    "Barney",
    "Beach break",
    "Bitchin'",
    "Bomb",
    "Blank",
    "Blown out",
    "Board",
    "Boardies",
    "Bottom turn",
    "Bogging",
    "Break",
    "Brah",
    "Bro",
    "Burn",
    "Carve",
    "Caught inside",
    "Channel",
    "Charging",
    "Cheater-five",
    "Cheehoo",
    "Chop",
    "Choppy",
    "Clam dragger",
    "Clean",
    "Closeout",
    "Cowabunga",
    "Crest",
    "Cross-step",
    "Cruise",
    "Curl",
    "Cutback",
    "Dawn patrol",
    "Deck",
    "Dick dragger",
    "Ding",
    "Drop",
    "Drop-in",
    "Drop-knee",
    "Duck diving",
    "Dude",
    "Face",
    "Fade",
    "Fin",
    "Fish",
    "Fish tail",
    "Firing",
    "Flat",
    "Floater",
    "Foam",
    "Frontside",
    "Froth",
    "Funboard",
    "Glassy",
    "Grabbing rail",
    "Greenroom",
    "Gnarly",
    "Good vibes",
    "Goofy footed",
    "Grom",
    "Gun",
    "Hang-five",
    "Hang-loose",
    "Hang-ten",
    "Haole",
    "Hawaiian",
    "Head-dip",
    "Heavy",
    "Heel-toe",
    "Hit the lip",
    "Hollow",
    "Inside",
    "Kahuna",
    "Kick-out",
    "Kook",
    "Layback",
    "Leash",
    "Left",
    "Leg rope",
    "Leggie",
    "Line-up",
    "Lines",
    "Lip",
    "Local",
    "Localism",
    "Log",
    "Longboard",
    "Lull",
    "Making-the-drop",
    "Men in the grey suits",
    "Messy",
    "Mini simmons",
    "Mush",
    "Mushy",
    "Noodle arms",
    "Nose",
    "Nose riding",
    "Nug",
    "Off the lip",
    "Off the wall",
    "Offshore",
    "Onshore",
    "Outside",
    "Outside break",
    "Out the back",
    "Over the falls",
    "Overhead",
    "Paddle",
    "Paddle battle",
    "Party wave",
    "Peak",
    "Pearl",
    "Peeling",
    "Pig dog",
    "Pintail",
    "Pit",
    "Pitted",
    "Pocket",
    "Pointbreak",
    "Poo-stance",
    "Psyched",
    "Pop up",
    "Pull-in",
    "Pumping",
    "Quad",
    "Quiver",
    "Rad",
    "Radical",
    "Rails",
    "Rag-dolled",
    "Rashguard",
    "Re-entry",
    "Reef break",
    "Regular footed",
    "Right",
    "Rip",
    "Rip current",
    "Riptide",
    "Rocker",
    "Rogue wave",
    "Salty",
    "Score",
    "Section",
    "Set",
    "Shacked",
    "Shaka",
    "Shape",
    "Shaper",
    "Shore",
    "Shore break",
    "Shortboard",
    "Shoulder",
    "Shred",
    "Shubie",
    "Sick",
    "Single fin",
    "Sketchy",
    "Skunked",
    "Slab",
    "Slotted",
    "Snake",
    "Soft top",
    "Soul-surfer",
    "Soul-arch",
    "Soup",
    "Spat out",
    "Spit",
    "Sponge",
    "Springsuit",
    "Stall",
    "Stance",
    "Stick",
    "Stinkbug",
    "Stoked",
    "Stringer",
    "Sucking dry",
    "SUP",
    "Surf",
    "Surf camp",
    "Surfboard",
    "Swell",
    "Switch",
    "Tail",
    "Take-off",
    "Throwing",
    "Thruster",
    "Tombstone",
    "Tow in",
    "Trough",
    "Tube",
    "Tubular",
    "Turtle roll",
    "Twin fin",
    "Twinnie",
    "Undertow",
    "Wahpah",
    "Wall",
    "Washing machine",
    "Wave",
    "Wax",
    "Wetsuit",
    "Whitewater",
    "Wipeout",
    "Worked",
    "Yeww",
    // Surf spots
    "Bells beach",
    "Blacks",
    "Cloud 9",
    "Jaws",
    "Malibu",
    "Mavericks",
    "North shore",
    "Pipeline",
    "Swamis",
    "The Wedge",
    // Surf brands
    "Billabong",
    "Body Glove",
    "Channel islands",
    "Dakine",
    "Former",
    "HaydenShapes",
    "Quicksilver",
    "Katin",
    "Matuse",
    "Mayhem",
    "NEEDessentials",
    "O'neill",
    "Patagonia",
    "Pyzel",
    "Rip Curl",
    "Roxy",
    "Rusty",
    "Rage",
    "Vans",
    "Vissla",
    "Xcel",
  }

  return wordList[rand.Intn(len(wordList))]
}
