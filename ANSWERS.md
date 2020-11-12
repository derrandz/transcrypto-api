#### Answers to PART 2 of assignment

##### What are, according to you, the three most underrated language features and why?

Question #1)

What are, according to you, the three most underrated language features and why? How can
Go still shoot you in the foot?

- This is very arguable, I am not sure if those are most underrated, but I can say that I find to be the best:
    - Channels:
        The way that channels have simplified the inter-processes communications is incredible
        I haven't made the time to implement something with channels in this repository, but I would have done so
        to (for instance) implement a little server that listens for system signals and communicates them through a channel
        for me to able to properly react to SIG_INT or something of the sort.
    
    - The Standard library:
        The standard library is very powerful, and the amount of useful interfaces that exposes for me to implement is incredible
        I was able to make an HTTP server from scratch using only net/http without resorting to anything else, and with only that
        and enough time, I suppose one can achieve anything

    - The linter: It's pretty smart and cool, and it comes out of the box. In comparison to the world of NodeJS where you spent 1/5 of your time
        configuring your toolchain, Golang's toolchain is just awesome as I have 0 work to do except for preparing a Makefile
    
    The list can go on but I would suffice with those for now

Question #2)

In your opinion, which cases justify using Unsafe? Explain your points

- The only circumstance that I can allow myself to use Unsafe is when I have made my data very complex to comply with Unix principles
    and transforming my data from a type to another and massaging it proves to be hasselous or causes my software to become very complex
    casting from an empty interface to a subtype to this to that.
    I'll access the data pointer unsafely to massage it however I want as I already have an idea of what I am receiving

    The other case would be some low-level memory manipulation, say, in a case where I am implementing some thread management library
    and I want to provide a memory-sharing feature with specific fast benchmarks, I might allow myself the unsafety of Unsafe, but it sure
    might come with a high price.