# 🌊 surfSum CLI 🏄
Lorem ipsum in surf slang 🤙 

Check out the [browser version here](https://gesty.dev/surfSum)

Inspired by [hipsum](https://hipsum.co) - Hipster Ipsum

## Usage
To run:
```bash
./surfSum-CLI
```
### Optional Flags
| Flags / Options | | Description |
| :-: | --- | ---|
| -sb | | adds "Shaka Brah!" to the start |
| -l | | mixes latin words in with the surf lingo |
| -p=<a number\> | | number of paragraphs to return (default 1, max 10) |
| -h or --help | | show help menu |

## Examples
**Command**: `./surfSum-CLI -sb`
```
Shaka Brah! Stick backside gun chop, lull patagonia overhead pearl spit. 
Aloha tow in tombstone alaia kick-out outside break, pitted twinnie. 
A-frame washing machine inside aggro pointbreak switch air shore break hang-ten. 
Hang-loose kick-out goofy footed cross-step bottom turn cheater-five, line-up 
carve. Pig dog chop shred off the wall billabong cheehoo greenroom.
```
**Command**: `./surfSum-CLI -l -sb`
```
Shaka Brah! Crest odio varius pumping, leggie floater congue, magna. Swamis 
cutback sit sponge, enim accumsan posuere. Cross-step shortboard fermentum 
quiver single fin worked. Wax rogue wave carve morbi dakine charging, rip 
montes. Dude washing machine surf pretium soul-surfer greenroom, twinnie 
springsuit. Patagonia pull-in brah pearl cutback egestas cloud 9 rusty. Reef 
break tellus closeout ridiculus, maecenas eiusmod bitchin' shaka nisl senectus. 
Slab foam xcel cross-step pumping do rhoncus.
```
**Command**: `./surfSum-CLI -l -p=2`
```
Haole the wedge reef break channel islands eget drop-knee, ultrices. Yeww alaia 
a-frame malesuada yeww take-off eget. Wahpah layback dolore patagonia, skunked 
billabong grom ding. Peak porttitor sick kahuna porttitor quicksilver phasellus 
lacinia. Fade shaka varius floater deck scelerisque neque ac, cowabunga. Glassy 
rusty blown out fish auctor floater. Springsuit lull stringer billabong twin 
fin rag-dolled, ankle slappers matuse greenroom layback. 

Pop up shacked facilisis hawaiian, localism suspendisse bomb tubular surf. Rails 
mush shaka stall channel islands dictum greenroom incididunt, patagonia. SUP 
turtle roll pharetra dolor alaia beach break funboard quisque. Drop-knee 
frontside springsuit curl, swamis ac urna eu. Kook convallis peak sit rad grom 
chop drop-in men in the grey suits.
```

<!-- TODO:
    - rewrite in js/html (new/separate github repo)
-->

<!-- Done:
    - option to start with "Shaka Brah" (flag)
    - separate file for sentence/paragraph building
    - variable sentence lengths
    - capitalize first letter of sentences
        - capitalized wordlist; `strings.ToLower` on non-first word
    - random commas to resemble real sentences
        - every few words (`i % 3 == 0`)
        - chance to insert comma (`rand.Intn(3)` 30% chance, `+ ","`)
    - variable paragraph lengths
    - option for returned paragraphs (flag)
    - show help/available options if no flags passed
    - option to mix in latin lorem ipsum (flag)
        - add file for latin words
    - add usage instructions/cheatsheet to README.md
-->
