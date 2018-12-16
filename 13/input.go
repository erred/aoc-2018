package main

var input = `
                          /------------------------------\                                           /----------\                                     
                          |             /----------------+----------------------------------------\  |          |                                     
                          |             |                |                                    /---+--+----------+-------------------------------\     
                          |             |        /-------+------------------------------------+---+--+----------+-----------------------------\ |     
                          |             |        | /-----+-----------------------------\/-----+---+--+----------+--------\                    | |     
    /-------------->\     |             |        | |     |                             ||     |   |  |   /------+--------+--------------------+\|     
    |               |     |             |   /----+-+-----+---------------\             ||     |   |  |   |      |        |                    |||     
    |    /----------+-----+\  /---------+---+----+-+-----+---------------+---------\   ||/----+---+\ |   |      |        |                    |||     
/---+----+----------+-----++--+----\    |   |    | |     |           /---+----\ /--+---+++----+---++-+---+-->->-+--------+--\                 |||     
|   |    |          |    /++--+----+----+---+----+-+-----+-----------+---+----+-+--+---+++----+---++-+---+------+--\     |  |                 |||     
|   |/---+----------+----+++--+----+--\ |   |    | |     |           |   |    | |  |   |||    |   || \---+------/  |     |  |                 |||     
|   ||   |          |    |||  |    |  |/+---+----+-+-----+-----------+--\|    | |  |   |||    |   ||     |         |     |  |          /------+++----\
|   ||   |     /----+----+++--+----+--+++---+----+-+-----+-----------+--++----+-+--+---+++----+---++-----+-------\ |     |  |          |      |||    |
|   ||   |/----+----+----+++--+----+--+++---+----+-+-----+-----------+--++-\  | |  |   |||/---+---++-----+--\    | |     | /+---------\|      |||    |
|   ||   ||    |    |    |||  |  /-+--+++---+----+-+-\   |/----------+--++-+--+-+--+---++++---+---++-----+--+----+-+-----+-++\        ||      |||    |
|   ||   ||  /-+----+----+++--+--+-+--+++---+----+-+-+---++----\  /--+--++-+--+-+--+---++++---+<--++-----+--+----+-+-----+-+++-------\||      |||    |
|   ||   ||  | |    |    ||| /+--+-+--+++---+----+-+-+---++----+--+--+--++-+--+-+--+---++++---+---++-----+--+----+-+-<---+-+++---\ /-+++------+++\   |
|   ||   ||  | |    |    ||| ||  | |  |||   |    | | |   ||    |  | /+--++-+--+-+--+---++++---+\  ||     \--+----+-+-----+-+++---+-+-+++------+/||   |
|   ||   ||  | |    |    ||| ||  | |  |||   |    | | |   ||    |  | ||  || |  |/+--+---++++---++--++---\    |    | |     | |||   | | |||      | ||   |
|   ||/--++--+-+----+----+++-++--+-+--+++---+----+-+-+---++----+--+-++--++-+--+++--+\  ||||   ||  ||   |    |    | |     | |||   | | |||      | ||   |
|   |||  ||  | |    |    ||| ||  | | /+++---+----+-+-+---++----+--+-++--++-+--+++--++--++++---++--++---+----+----+-+---\ | |||   | | |||      | ||   |
|   |||  ||  | |    |    ||| ||  | | ||||   |    | | |   ||    |  | ||  || |  |||  || /++++---++--++---+----+----+-+---+-+-+++---+-+-+++----\ | ||   |
|   |||  ||  | |    |    ||| ||  | | |||\---+----+-+-+-->++----+--+-++--++-+--+++--++-+++++---++--/|   |    |    | |   | | |||   | | |||    | | ||   |
|   |||  ||  | |    |    ||| ||  | | |||    |    | | |   ||    |  | || /++-+--+++--++-+++++---++---+---+----+----+-+---+-+-+++---+\| |||    | | ||   |
|   |||  ||  \-+----+----+++-++--+-+-+++----+----+-+-+---++----/  | || ||| |  |||  || |||||   ||   |   |    |    | |   | | |||   ||| |||    | | ||   |
|   |||  || /--+----+----+++-++--+-+-+++----+--\ | | |   ||/------+-++-+++-+--+++--++-+++++---++---+---+----+----+-+-\ | | |||   ||| |||    | | ||   |
|  /+++--++-+--+----+----+++-++--+-+-+++----+--+-+-+-+---+++------+-++-+++-+--+++--++\|||||   ||   |   |    |    | | | | | |||   ||| |||    | | ||   |
|  ||||  || | /+----+----+++-++--+-+-+++----+--+-+-+-+---+++------+-++-+++-+--+++--++++++++---++---+---+----+----+-+-+-+-+\|||   ||| |||    | | ||   |
|  ||\+--++-+-++----+----+++-++--+-+-+/|    |  | | | |   |||      |/++-+++-+--+++--++++++++---++---+---+----+\   | | | | |||||   ||| |||    | | ||   |
|  || |  || | || /--+--<-+++-++--+-+-+-+----+--+-+-+-+---+++------++++-+++-+--+++--++++++++---++---+--\|    ||   | | | | |||||   ||| |||    | | ||   |
| /++-+--++-+-++-+--+----+++-++--+-+-+-+----+--+-+-+\|   |||    /-++++-+++-+--+++--++++++++---++---+--++---\||   | | | | |||||   ||| |||    | | ||   |
| ||| |  || | || |  |    ||| ||  | | | |    |  | | |||   |||/---+-++++-+++-+--+++--++++++++---++-\ |  ||   |||   | | | | |||||   ||| |||    | | ||   |
| ||| |  || |/++-+--+----+++-++--+-+-+-+----+--+-+-+++---++++---+-++++-+++-+--+++\ |||||||\---++-+-+--++---+/|   | | | | |||||   ||| |||    | | ||   |
| ||| |  || |||| |  |    ||| ||  | | |/+----+--+-+-+++---++++---+-++++-+++-+--++++-+++++++----++-+-+--++---+-+---+-+-+-+-+++++\  ||| |||    | | ||   |
| ||| |  || |||| |  |    ||| ||  | | |||    |  | | |||   ||||   | |||| ||| |  |||| |||||||    || | |  ||   | |   | | | | ||||||  ||| |||    | | ||   |
| ||| |  || |||| |  |    |\+>++--+-+-+++----+--+-+-+++---/|||   | |||| ||| |  |||| |||||||    || | |  ||   | |   | | | | ||||||  ||| |||    | | ||   |
| ||| |  || |||| |  |   /+-+-++--+-+-+++----+--+-+-+++--\ |||   | |||| ||| |  |||| |||||||    || | |/-++---+-+---+-+-+-+-++++++--+++-+++---\| | ||   |
| ||| |  || ||||/+--+---++-+-++--+\| ||\----+--+-+-+++--+-+++---+-++++-+/| |  |||| |||||||    || | || ||   | |   | | | | ||||||  ||| |||   || | ||   |
| ||| |  \+-++++++->+---++-/ ||  \++-++-----+--+-+-++/  | |||   | |||| | | |  |||| |||||||    || |/++-++---+-+---+-+-+-+-++++++--+++-+++\  || | ||   |
| ||| |   | ||||||/-+---++---++---++-++-----+--+-+-++---+-+++---+-++++-+-+-+-\|||| |||||||    || |||| ||   | |   | | | | ||||||  ||| ||||  || | ||   |
| ||| |   | ||||||| |   ||   ||   || || /---+--+-+-++---+-+++---+-++++-+-+-+-+++++-+++++++----++-++++-++---+-+\  | | | | ||||||  ||| ||||  || | ||   |
| ||| |   | ||||||| |   ||   ||   || || |   |  | | ||   |/+++---+-++++-+-+-+\||||| |||||||    || |||| ||   | ||  | | | | ||||||  ||\-++++--++-+-+/   |
| ||| |   | ||||||| |   ||   ||   || || |   |  | | ||   |||||/--+-++++-+-+-+++++++-+++++++----++-++++-++---+-++--+-+-+-+-++++++--++\ ||||  || | |    |
| ||| |   | ||||||| |   ||   ||   || || |   |  | | ||  /++++++--+-++++-+-+-+++++++-+++++++----++-++++-++---+-++--+-+-+-+-++++++--+++-++++\ || | |    |
| ||| |   | ||||||| |   ||   \+---++-++-+---+--+-+-++--+++++++--+-++++-+-+-+++++++-+++++++----++-++++-++---+-++--+-+-+-+-++++++--/|| ||||| || | |    |
| ||| |   | ||||||| |   ||  /-+---++-++-+---+--+-+-++--+++++++--+-++++-+-+-+++++++-+++++++----++-++++-++---+-++--+-+-+\| ||||||   || ||||| || | |    |
| ||| |   | ||||||| |   ||  | |   || || |   |  | | ||  |||||||  | |||| | | ||||||| |||||||    || |||| ||   | ||  | | ||| ||||||   || ||||| || | |    |
| ||| |   | ||||||| |   ||  | |   || || |   |  | | ||  |||||||  | |||| | | ||||||| |||||||    || ||||/++---+-++--+-+-+++>++++++---++-+++++-++-+\|    |
\-+++-+---+-+++++++-+---++--+-+---+/ || |   |  | | \+--+++++++--+-++++-+-+-+++++++-++++/||    || |||||||   | ||  | | ||| ||||||   || ||^|| || |||    |
  ||| |   | ||||||| |   || /+-+---+--++-+---+--+-+--+--+++++++-\| |||| | | ||||||| |||| ||    || |||||||   | ||  | | ||| ||\+++---++-+/||| || |||    |
  ||| |   | ||||||| |   || || |   |  || |   |  | |  |  ||||||| || |||\-+-+-+++/||| |||| ||    || |||||||   | ||  | | ||| || |||   || | ||| || |||    |
  ||| |   | ||||||| |   || || |   |  || |   |  | |  |  ||||||| || |||  | | ||| ||| |||| ||/---++-+++++++---+-++-\| | ||| || |||   || | ||| || |||    |
  ||| |   | ||||||| |   ||/++-+---+--++-+---+--+-+--+--+++++++-++-+++--+-+-+++-+++-++++-+++---++-+++++++---+-++-++-+-+++-++-+++\  || | ||| || |||    |
  ||| |   | ||||||| |   |\+++-+---+--++-+---+--+-+--+--+++++++-++-+++--+-+-+++-+++-++++-+++---++-+++++++---+-++-++-/ ||| || ||||  || | ||| || |||    |
  \++-+---+-+++++++-+---+-+++-+---+--++-+---+--+-+--/  ||||\++-++-+++--+-+-+++-+++-++++-+++---++-+++++++---+-++-++---/|| || ||||  || |/+++-++-+++---\|
   || |   | ||||||| |   | ||| | /-+--++-+---+--+-+-----++++-++-++-+++--+-+-+++-+++-++++-+++---++-+++++++---+-++-++----++-++-++++--++-+++++\|| |||   ||
   || |   \-+++++++-+---+-+++-+-+-+--++-+---+--+-+-----++++-++-++-+++--+-+-/|| ||| |||| \++---++-+++++++---+-++-++----++-/| ||||  || |||||||| |||   ||
   || |     ||||||| | /-+-+++-+-+-+--++-+-\ |/-+-+-----++++-++\|| |||  | |  || \++-++++--++---++-++++++/   | || ||    ||  | ||||  || |||||||| |||   ||
   || |     |||||\+-+-+-+-+++-+-+-+--++-+-+-++-+-+-----++++-+++++-+++--+-+--++--++-++++--++---++-+++++/   /+-++-++-\  ||  | ||||  || |||||||| |||   ||
   || |     ||||| | | | | ||| | | |  || | | || | |     |||| ||||| |||  | |/-++--++-++++--++---++-+++++----++-++-++-+--++--+-++++--++-++++++++\|||   ||
  /++-+-----+++++-+-+-+-+-+++-+-+-+--++-+-+-++-+-+-----++++-+++++-+++--+-++-++--++-++++--++---++-+++++--\ || || || |  ||  | ||||  || ||||||||||||   ||
  ||| |     ||||| | | | | ||| | | |  || | | || | |     |||| ||||| |||  | || ||  || ||||  ||   || |||||  | || || || |  ||  | ||||  || ||||||||||||   ||
  ||| |    /+++++-+-+-+-+-+++-+-+-+--++-+-+-++-+-+-----++++-+++++-+++--+-++-++\ || ||||  ||   || |||||  | || || || |  ||  | ||||  || ||||||||||||   ||
  ||\-+----++++++-+-/ | | ||| | | |  |\-+-+-++-+-+-----++++-+++++-+++--+-++-+++-++-++++--++---++-+++++--+-++-++-++-+--++--+-++/|  || ||||||||||||   ||
  ||  |    |||||| |   | | ||| | | |  |  |/+-++-+-+----\|||| ||||| ||\--+-++-+++-++-++++--++---+/ ||||\--+-++-++-++-+--++--+-++-+--++-++++++++++/|   ||
  ||/-+----++++++-+---+-+-+++-+-+-+--+--+++-++\| |    |||||/+++++-++---+-++-+++-++-++++--++---+--++++---+-++-++-++-+--++--+-++-+--++\|||||||||| |   ||
  ||| |    |||||| |   | | ||| | | |  |  ||| |||| |    ||||||||||| ||   | || ||| || ||||  ||   |  ||||   | || || || |  ||  | || |  ||||||||||||| |   ||
  ||| |    |||||| |   | | |||/+-+-+--+--+++-++++-+----+++++++++++-++---+-++-+++-++-++++--++---+--++++---+-++-++-++-+\ ||  | || |  ||||||||||||| |   ||
  ||| | /--++++++-+---+-+-+++++-+-+--+--+++-++++-+----+++++++++++-++---+-++-+++-++-++++--++---+--++++---+-++\|| || || ||  | || |  ||||||||||||| |   ||
  ||| | |/-++++++-+---+-+-+++++-+-+--+-\||| |||| |    ||||||||||| ||   | || ||| || ||||  ||   |  ||||   | ||||| || || ||  | || |  ||||||||||||| |   ||
  ||| | || |||||| |   | | ||||| \-+--+-++++-++++-+----+++++++++++-++---+-++-+++-++-++++--++---+--++++---+-+++++-++-++-++--+-++-+--++++++++/|||| |   ||
  ||| | || |||||| |   | | ||||\---+--+-++++-++++-+----+++++++++++-++---+-++-+++-++-/|||  ||   |  ||||   | ||||| || || ||  | || |  |||||||| |||| |   ||
/-+++-+-++-++++++-+---+-+-++++----+--+-++++-++++-+-\/-+++++++++++-++---+-++-+++-++--+++--++---+\ ||||   | ||||| || || ||  | || |  |||||||| |||| |   ||
| ||| | || |||||| |   | | ||||    |  | |||| |||| | || ||||||\++++-++---+-++-+++-++--+++--++---++-/|||   | ||||| || || ||  | || |  |||||||| |||| |   ||
| ||| | || |||||| |   | | ||||    |  | |||| |||| | || |||||| |||| ||  /+-++-+++-++--+++--++---++--+++---+-+++++-++-++-++--+-++-+--++++++++-++++-+\  ||
| ||| | || |||||| |   | | ||||    |  | |||| |||| \-++-++++++-++++-++--++-++-+++-++--+++--++---++--+++---+-+++++-++-++-++--+-++-+--++++++++-+++/ ||  ||
| ||| | || |||||| |   | |/++++----+--+-++++-++++---++-++++++-++++-++--++-++-+++-++--+++--++-\ ||  |||   | ||||| || || ||  | || |  |||||||| |||  ||  ||
| ||| | || |||||| |   | ||||||    |  | |||| ||||   || |||||| |||| ||  || || ||| ||  |||  || | ||  |||   | ||||| || || ||  | || |  |||||||| |||  ||  ||
| ||| | ||/++++++-+--\| ||||||    |  | |||| ||||   || |||||| |||| ||  || || ||| ||  |||  \+-+-++--+/\---+-+++++-++-++-++--+-++-+--++++++++-/||  ||  ||
|/+++-+-+++++++++-+--++-++++++\   |  | |||| ||||   || |||||| |||| ||  |\-++-+++-++--+++---+-+-++--+-----+-+++++-++-++-++--+-++-+--/|||||||  ||  ||  ||
||||| | ||||||||| |  || |||||||   |  | |||| ||||   || |||||| |||| ||  |  || ||| ||  |||   | | ||  |     | ||||| || || ||  | || |   |||||||  ||  ||  ||
||||| | ||||||||| |  || |||||||   |  | |||| ||||   || |||||| |||| ||  |  || ||| ||  |||   | | ||  |     | ||||| || || ||  | || |   |||||||  ||  ||  ||
||||| | ||||||||| |  || |||||||   |  | |||| ||||   || |||||| |||\-++--+--++-+++-++--+++---+-+-++--+-----+-+/||| || || ||  | || |   |||\+++--++--++--/|
||||| | ||||||||| |  || |||||||   |  | |||| ||||   || |||||| |||  ||  |  || ||| ||  |||   | | ||  |     | | ||| || || ||  | || |   ||| |||  ||  ||   |
||||\-+-+++++++++-+--++-+++++++---+--+-++++-++/|/--++-++++++-+++--++--+--++-+++-++--+++---+-+-++--+-----+-+-+++-++-++-++\ | || |   ||| |||  ||  ||   |
||||  | ||||||||| |  || |||||||   |  | |||| || ||  || |||||| |||  ||  |  || ||| \+--+++---+-+-++--+-----+-+-+++-++-++-+++-+-/| |   ||| |||  ||  ||   |
||||  | ||||||||| |  || |||||||/--+--+-++++-++-++--++-++++++-+++--++--+--++-+++--+--+++---+-+-++--+-----+-+-+++\|| || ||| |  | |   ||| |||  ||  ||   |
||||  | ||||||||| |  ||/++++++++--+--+-++++-++-++--++-++++++-+++--++--+--++-+++--+--+++---+-+-++--+-----+\| |||||| || ||| |  | |   ||| |||  ||  ||   |
||||  | |||\+++++-+--+++++++++++--+--+-++++-++-++--++-++++++-+++--++--+--++-++/  |  |||   | | ||  |     ||| |||||| || ||| |  | |   ||| |||  ||  ||   |
||||  | ||| ||||| |  |||\+++++++--+--+-++++-++-++--++-++/||\-+++--++--+--++-++---+--+++---+-+-++--+-----+++-++++++-++-+++-+--+-+---+/| |||  ||  ||   |
||||  |/+++-+++++-+--+++-+++++++--+--+-++++-++-++--++-++-++--+++--++--+--++-++---+--+++---+-+-++--+---\ ||| |||||| || ||| |  | |   | | |||  ||  ||   |
||||  ||||| ||||| |  ||| |||||||  |  | |||| || ||  || ||/++--+++--++--+--++-++---+--+++---+-+-++--+---+-+++-++++++-++-+++-+--+-+---+-+-+++--++--++-\ |
||||  ||||| \++++-+--+++-+++++++--+--+-++++-++-/|  || |||||  |||  ||  |  || ||   |  |||   | | \+--+---+-+++-++++++-++-+++-+--+-+---+-+-+++--++--/| | |
||||  |||||  |||| |  ||| |||||||  |  | |||| ||  |/-++-+++++--+++--++--+--++-++---+--+++---+-+--+--+---+-+++-++++++\|| ||| |  | |   | | |||  ||   | | |
||||  |||||  |||| |  ||| |||||||  |  | |||| ||  || || |||||  |||  ||  \--++-++---+--+++---+-+--+--+---+-+++-+++++++++-+++-+--+-+---+-+-+++--++---/ | |
||||  |||||  |||| |  ||| |||||||  |  | |||| ||  || || |||||  |||  |\-----++-++---+--+++---+-+--+--+---+-+++-+/||||||| ||| |  | |   | | |||  ||     | |
||||  |||||/-++++-+--+++-+++++++--+--+-++++-++--++-++-+++++--+++--+------++-++--\|  |||   | |  |  |   | ||| | ||||||| ||| |  | |   | | \++--++-----+-/
||||  |||||| |||| |  ||| |||||||  |  | |||| ||  || || |||||  |||  |      || ||  ||  |||   | |  |  |   | ||| | ||||||| ||| |  | |   | |  ||  ||     |  
||||  |||||| |||| | /+++-+++++++--+--+-++++-++--++-++-+++++--+++--+------++-++--++--+++---+-+--+--+--\| ||| | ||||||| ||| |  | |   | |  ||  ||     |  
||||  |||||| |||| | |||| |||||||  |  | |||| ||  || || |||||  |||  |      || ||  ||  |||   | |  |  |  || ||| | ||||||| ||| |  | |   | |  ||  ||     |  
||||  |||||| |||| | |||| |||||||  |  | |||| ||  || || |||||  |||  \------++-++--++--+++---+-+--+--+--++-+++-+-+++++++-+++-+--+-+---+-/  ||  ||     |  
||||  |||||| |||| | |||| ||\++++--+--+-++++-++--++-++-+++++--++/         || ||  ||  |||   |/+--+--+--++-+++\| ||||||| ||| |  |/+---+---\||  ||     |  
||||  |||||| |||| | |||| || ||||  |  | |||| \+--++-++-+++++--++----------/| ||  ||  |||   |||  |  |  || ||||| ||||||| ||| |  |||   |   |||  ||     |  
||||  ||||\+-++++-+-+/|| |^ ||||  |  | ||||/-+--++-++-+++++--++-----------+-++--++--+++---+++--+--+--++-+++++-+++++++-+++-+--+++---+---+++--++---\ |  
||||  |||| | ||||/+-+-++-++-++++-\|  | ||||| |  || || ||||\--++-----------+-++--++--+++---+++--+--+--++-+++++-+++++++-+++-+--/||   |   |||  ||   | |  
||||  |||| | ||\+++-+-++-++-++++-++--+-+++++-+--++-++-++++---++-----------+-++--++--+++---+++--+--+--++-+++++-+++/||| ||| |   ||   |   |||  ||   | |  
||||  ||\+-+-++-+++-+-++-++-++++-++--+-+++++-+--++-++-++++---++-----------+-++--++--+++---+++--+--+--++-++++/ ||| ||| ||| |   ||   |   |||  ||   | |  
||||  || | | || ||| | || || |||| ||  | ||||| |  || || ||||   ||           | ||  ||  ||\---+++--+--+--++-++++--+++-+++-+++-+---++---+---+++--/|   | |  
||||  || | | || ||| | || |\-++++-++--+-+++++-+--++-++-++++---++-----------+-++--++--++----+++--+--+--++-++++--+++-+++-+++-+---+/   |   |||   |   | |  
||||  || | | || ||| | || |  |||| ||  | ||||| |  || || ||||   ||/----------+-++--++--++----+++--+--+--++-++++--+++-+++-+++-+---+----+---+++---+---+-+-\
||||  || | | || ||| | || |  |||| ||  | |\+++-+--++-++-++++---+++----------+-++--++--++----+++--+--+--++-++++--/|| ||| ||| |   |    |   |||   |   | | |
||||  || | | || ||| \-++-+--++++-++--+-+-+++-+--++-++-++++---+++----------+-++--++--++----+++--+--+--/| ||||   || ||| ||| |   |    |   |||   |   | | |
||||  \+-+-+-++-+++---++-+--++++-++--+-+-+++-+--++-++-++++---+++----------+-++--++--/|    |||  |  |   | ||||   || ||| ||| |   |    |   |||   |   | | |
||||   | | | || |||   || |  |\++-++--+-+-+++-+--++-++-++++---+++----------+-++--++---+----+++--+--+---+-++++---++-++/ ||| |   |    |   |||   |   | | |
||||   | | | || |\+---++-+--+-++-/|  \-+-+++-+--++-++-++++---+++----------+-++--++---+----+++--+--+---+-++++---++-++--+/| |   |    |   |||   |   | | |
\+++---+-+-+-++-+-+---++-+--+-++--+----+-+++-+--++-/| |||\---+++----------+-/|  ||   |    \++--+--+---+-++++---+/ ||  | | |   |    |   |||   |   | | |
 |||/--+-+-+-++-+-+---++-+->+-++--+-\  | ||| |  ||  | |||    \++----------+--+--++---+-----++--+--+---+-++++---+--++--+-+-+---+----/   |||   |   | | |
 ||||  | | | || | |   || |  | ||  | |  | ||| |  ||  | |^|     ||          |  |  ||   |     \+--+--+---+-+++/   |  ||  | | |   |        |||   |   | | |
 ||||  | | | |\-+-+---++-+--+-++--+-+--+-+++-+--++--+-+++-----++----------+--+--++---+------+--+--+---+-+++----+--++--+-+-/   |        |||   |   | | |
 ||||  | | | |  | |   || |  | ||  | |  | ||| |  ||  | |||     ||         /+--+--++---+------+--+--+---+-+++----+--++--+-+-----+-\      |||   |   | | |
 ||||  | | | |  | \---++-+--+-++--+-+--+-+++-+--++--+-+++-----++---------++--/  ||   |      |  |  |   | |||    |  ||  | |     | |    /-+++---+\  | | |
 ||\+--+-+-+-+--+-----++-+--+-++--+-+--+-+++-+--++--+-+++-----++---------++-----++---/      |  |  |   | |||    |  ||  | |     | |    | |||   ||  | | |
 || |  | | | |  |     || |/-+-++--+\|  | ||| |  ||/-+-+++--\  ||         ||     ||          |  |  |   | |||    |  ||  | |     | |    | |||   ||  | | |
 || |  | | | |  |     || || | ||  |||  | |||/+--+++-+\|||  |  ||         ||     ||          |  |  |   | |||    |  ||  | |     | |    | |||   ||  | | |
 || |  | | | \--+-----++-++-+-++--+++--+-+++++--+++-+++++--+--++---------++-----+/          |  |  |   | |||    |  ||  | |     \-+----+-/||   ||  | | |
 || |  | | |    |     || || | ||  |||  | |||||  ||| |||||  |  ||         ||     |           |  |  |   | |||    |  ||  | |       |    |  ||   ||  | | |
 || |  | | |    |     \+-++-+-++--+++--+-+/|||  ||| |||||  |  ||         ||     |           |  |  |   | |||    |  ||  | |       |    |  ||   ||  | | |
 || |  | | |    |      | || | ||  |||  | | |||  ||| \++++--+--++---------++-----+-----------+--/  |   | |||    |  ||  | |       |    |  ||   ||  | | |
 || |  \-+-+----+------+-++-+-++--+++--+-+-+++--+++--++++--+--++---------++-----+-----------+-----+---/ |||    |  ||  | |       |    |  ||   ||  | | |
 || |    | |    |      | || | ||  |||  | \-+++--+++--+/|| /+--++---------++-----+-----------+-----+----\|||    |  ||  | |       |    |  ||   ||  | | |
 || |    | |    |      | || | ||  ||v  |   |||  |||  | || ||  ||         ||     |           |     |    ||||    |  ||  | |       |    |  ||   ||  | | |
 || |    | \----+------+-++-+-++--+++--+---+++--+++--+-++-++--++---------++-----/           |     |    ||||    |  ||  | |       |    |  ||   ||  | | |
 || |    |      ^      | || | ||  |||  |   |||  |||  | |\-++--++---------++-----------------+-----+----++++----+--++--+-+-------+----+--++---++--+-/ |
 || |    |      \------+-++-+-++--/||  |   |||  \++--+-+--++--++---------++-----------------+-----+----++++----+--++--+-/       |    |  ||   ||  |   |
 || |    |             | || | ||   ||  |   |||   ||  | |  ||  ||         ||                 |     |    ||||    |  ||  |         |    |  ||   ||  |   |
 || |    |             | \+-+-++---++--+---+++---++--+-+--++--++---------++-----------------/     |    ||||    |  ||  |         |    |  ||   ||  |   |
 || |    |             |  | \-++---++--+---+++---++--+-+--++--++---------++-----------------------+----++++----+--++--/         |    |  ||   ||  |   |
 \+-+----+-------------+--+---/|   ||  |   |||   ||  | |  ||  ||         ||                       \----++++----+--++------------+----+--/|   ||  |   |
  | |    |             |  |    \---++--+---+++---++--+-+--++--++---------++----------------------------++++----/  ||            |    |   |   ||  |   |
  \-+----+-------------+--+--------++--+---+++---++--+-+--++--++---------++----------------------------+/|\-------+/            |    |   |   ||  |   |
    |    |             |  |        ||  |   |||   ||  | |  \+--++---------++--------------------<-------/ |        |             |    |   |   ||  |   |
    |    |             |  |        ||  |   |||   \+--+-+---+--++---------++------------------------------+--------/             |    |   |   ||  |   |
    |    |             |  |        ||  |   |||    \--+-+---/  ||         ||                              |                      |    |   |   ||  |   |
    |    |             \--+--------++--+---+++-------+-+------++---------++------------------------------/                      |    |   |   ||  |   |
    \----+----------------+--------+/  |   |||       | |      ||         \+-----------------------------------------------------/    \---+---+/  |   |
         |                |        |   |   \++-------+-+------++----------+--------------------------------------------------------------+---+---/   |
         |                |        |   |    ||       | |      |\----------+--------------------------------------------------------------+---+-------/
         |                \--------/   |    |\-------+-+------/           |                                                              |   |        
         |                             |    |        | \------------------+--------------------------------------------------------------/   |        
         \-----------------------------/    \--------/                    \------------------------------------------------------------------/        
`