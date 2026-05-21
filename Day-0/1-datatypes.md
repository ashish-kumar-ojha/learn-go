In all the programming there are different types of data/variables, similarly ***Go*** has those types.

* ***int*** - stores integers (whole numbers), such as 3 or -300
* ***float32***- stores floating point numbers, with decimals, such as 11.99 or -199.99
* ***string*** - stores text, such as "Hello World". String values are surrounded by double quotes



## Declaring Variables

There are two ways to declare variable in ***Go***

1. using ***var*** keyword
2. using ***":="*** sign

#### Using ***var***

`var variableName type = value`

***Example***

`var student string = "nameHere"`

`var age int = 29`

As we have to mention type here, it’s kind of a hassle, right?
*using ":=" solves this*.

#### Using ***:=*** Sign

When we use this sign, the type of variable will be ***inferred*** from the value which means that the ***compiler*** decides the type based on the value

***Example***

`name := "NameHere"`

`age := 29`

`isActive := false`