# Programming Assignment

## Task description

Your task is to implement an interface for generating random texts based on
previously seen text corpus.

You should timebox this task to around 1 hour. We do not expect a perfect solution (or even complete necessarily). 
Instead we are interested in seeing how you approach the problem and understand the problem description. 

To accomplish this task, we want to parse the given text into a list of words.
These words are then structured into "trigrams" which are simply groups of three words.
Now for each such trigram we calculate the frequency of how often a word appears
after the previous two words. Using these frequencies we can generate text based
on the last two words that we have generated.

### Example

Lets assume our input text corpus is the sentence:

> To be or not to be that is the question

From this we may extract the words:

```
["To", "be", "or", "not", "to", "be", "that", "is", "the", "question"]
```

The eight trigrams are:

```
["To", "be", "or"]
["be", "or", "not"]
["or", "not", "to"]
["not", "to", "be"]
["to", "be", "that"]
["be", "that", "is"]
["that", "is", "the"]
["is", "the", "question"]
```

The frequencies count how often we have seen a word, given the previous two words.
For the input above we would end up with the following frequencies:

| Prefix  | word     | count |
|---------|----------|-------|
| be or   | not      | 1     |
| be that | is       | 1     |
| is the  | question | 1     |
| not to  | be       | 1     |
| or not  | to       | 1     |
| that is | the      | 1     |
| to be   | or       | 1     |
| to be   | that     | 1     |

Now when we want to generate a new text we first pick one of the trigrams at
random and then, word by word, we create the text using the calculated frequencies.

For instance we may start with "[To be or]" as the first trigram. From this we
take the first two words to start our text:

> To be

We now pick the next word from the two options we have for this prefix which are
"or" (1) and "that" (1). Both words have the same count so the chances to pick
either is a 50:50 (or 1:1). Lets say we pick "that":

> To be that

The new prefix now is created from the last two words of our generated text:
"be that". Looking at our frequency table again we see there is only a single
option we can take from here which is the word "is":

> To be that is

You can see with the limited corpus we provided there is not a lot of room to
generate an interesting text. However if we train our model on a large enough
corpus, we will have more options, often with different counts for the next
possible word. These counts should be used to make a _weighted_  random
selection of the next word. For instance if the options for the next word are
"a" with a weight of 2 and "b" with a weight of 1, the chance that we pick "a"
should be twice as high as the chance of picking "b".

## Instructions

To get you started we already provide a main function that reads in a text file
and parses it into a list of trigrams. You need to create a new type in the `solution` 
directory that implements the `Solution` interface that is specified right above the `main`
function in the `main.go` file. 

Additionally we provide some helper functions in the `trigram` package. You can
use those functions or completely ignore them if you do not need them. 

To play with the program that you are about to write, we also provide some example
text files in the `corpus` directory. Each file is a simple text file without any
styling or markdown. They should all be accepted by the parser that we implemented
for you.

The task is complete when we can execute it to generate a random text based on a
previously learned corpus file. If something is unclear you can always come and
ask us for advice. Additionally you are free to use the internet or any other
tooling you like. Please do however complete the program in Go.

If you want to run your code locally, using a default corpus to generate 100 words, run:

```
go run main.go
``` 

When you are done with the task, please zip or tar your files (or simply the
full directory) and send it to us via this Google form: https://forms.gle/y2GxnSNkMSavTKkK9 (or directly to the recruiter you were in contact with).

Good luck and happy hacking! \ʕ◔ϖ◔ʔ/ 

### A Word of Advice

Don't worry too much if you will complete the task completely in the time box. Instead
budget your time well to get together some code we may discuss in an interview. If some of your
code is not yet optimal we recommend you leave a comment and move on. You can always
later optimize it if you have more time or we can talk about it in the interview.

You will also find a function to normalize words in the provided `trigram` package.
At the beginning ignore this and try to complete a first implementation that generates
some text. When you are happy with the result and still have time you can think about
why and how you would use normalization.

### Bonus tasks

If you happen to have extra time here are some suggestions what you can add:

- implement tests for the solution (we expect most people will do this anyway)
- can you benchmark your code to show where it can be improved to be faster or use less memory?
- review and change our provided code
  - maybe the `Solution` interface could be improved?
  - have a look at the parser and see if it can be optimized
  - how would you test the parser code?
