# Assignment 5 - An Exercise in AI-Assisted Programming

summarize your experiences with each of the three methods: automated code generation, AI-assisted programming, and AI-generated code

For this assignment, I compared the performance of AI-assisted programming and AI-generated code to the performance of the code that I created for Assignment 4. In that assignment, I tested the performance of Go against Python and R to generate linear regression coefficients for the Anscombe Quartet. Ultimately, I found that the AI-Assisted programming performed the best, including over the code from Assignment 4. For this test, I repeated the Assignment 4 specifications, inlcuding that the package used for statistics generation.

| Language | Assignment 4 | AI-Assisted | AI-Generated |
|:--------:|:------------:|:-----------:|:------------:|
| Go | 0.39 sec | 0.40 sec | 0.45 sec |
| Python | 0.43 sec| 0.02 sec | 0.04 sec |
| R | 0.02 sec | 0.02 sec | 0.02 sec |

There were three methods of integrating AI with programming for this assignment: AI-assisted programming, AI-generated code, and automated code generation. For the AI-Assisted progamming component, I used the AI to review code I had already written. My goal was to improve functionality I was already sure worked, and the AI performed well. This method was especially good at generating test cases with edge cases that I hadn't thought of. The AI-generated code was less functional, and seemed to prefer for-loops in places where they were not necessary. It required more coaching, and a lot of time was spent on polishing a flawed solution. Automated code generation was a feature that supplamented both approaches, and I decided to turn it off. The attempted predictions were occasionally useful, but more often provided suggestions completely without context.

## Recommendation to Management
 
 **Does it produce better code?**
 AI-Assisted Programming produced functional code in all cases, though it did require significant human intervention to be brought to the same level as the human programming in Assignment 4. The Chat Agents had a particular blind-spot for writing idiomatic Go code, seeming to prefer syntax that was more familiar to Python or R scripts. The Python and R scripts performed better than the human-generated code, so Go may be problem-area for the LLM supporting the Chat Agent.

 **What are the possibilities for reducing the programming workload in the startup?**
 The main way that I would recommend the programming workload be reduced in the startup with respect to AI-Assisted programming is by generating tests and benchmarks for the code. Human generated code is more proficient a big-picture planning, but AI-assisted programming would be helpful with implementation, as long as the programmer also ensures that the code is idiomatic.

 What is your recommendation to the firm regarding its staffing of programmers and software engineers?

## Background
The purpose of the Anscombe Quartet is to demonstrate how it is important to visualize datasets instead of relying on pure data. Each of the datasets in the quartet returns identifical regression coefficients, despite being very different. 

For more information, see:
Anscombe, F. J. 1973, February. "Graphs in Statistical Analysis." *The American Statistician 27(1)*: 17–21. Available online at [https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf).


## Roles of Programs and Data
Note any training materials.

### AI_Assisted
text
### AI_Generated
text

## Application
text

## Use of AI
text

### Automated Code Generation
- Uses pre-defined rules and templates from high-level specifications
### AI-Assisted Programming
- Revise original code with assistance of Github Copilot?
- Ways the Github Copilot could assist with the assignment
- Different proficiency in each language (Go, Python, R)
### AI Generated Code
- Step-by-Step Review of Work
- Prompt review
- Describe how files were obtained


