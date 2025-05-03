# Assignment 5 - An Exercise in AI-Assisted Programming

For this assignment, I compared the performance of AI-assisted programming and AI-generated code to the performance of the code that I created for Assignment 4. In that assignment, I tested the performance of Go against Python and R to generate linear regression coefficients for the Anscombe Quartet. Ultimately, I found that the AI-Assisted programming performed the best, including over the code from Assignment 4. For this test, I repeated the Assignment 4 specifications, inlcuding that the package used for statistics generation.

| Language | Assignment 4 | AI-Assisted | AI-Generated |
|:--------:|:------------:|:-----------:|:------------:|
| Go | 0.39 sec | 0.40 sec | 0.45 sec |
| Python | 0.43 sec| 0.02 sec | 0.04 sec |
| R | 0.02 sec | 0.02 sec | 0.02 sec |

There were three methods of integrating AI with programming for this assignment: AI-assisted programming, AI-generated code, and automated code generation. For the AI-Assisted progamming component, I used the AI to review code I had already written. My goal was to improve functionality I was already sure worked, and the AI performed well. This method was especially good at generating test cases with edge cases that I hadn't thought of. The AI-generated code was less functional, and seemed to prefer for-loops in places where they were not necessary. It required more coaching, and a lot of time was spent on polishing a flawed solution. Automated code generation was a feature that supplamented both approaches, and I decided to turn it off. The attempted predictions were occasionally useful, but more often provided suggestions completely without context.

## Recommendation to Management
 
 **Does it produce better code?**
 <br>
 AI-Assisted Programming produced functional code in all cases, though it did require significant human intervention to be brought to the same level as the human programming in Assignment 4. The Chat Agents had a particular blind-spot for writing idiomatic Go code, seeming to prefer syntax that was more familiar to Python or R scripts. The Python and R scripts performed better than the human-generated code, so Go may be problem-area for the LLM supporting the Chat Agent.

 **What are the possibilities for reducing the programming workload in the startup?**
 <br>
 The main way that I would recommend the programming workload be reduced in the startup with respect to AI-Assisted programming is by generating tests and benchmarks for the code. Human generated code is more proficient a big-picture planning, but AI-assisted programming would be helpful with implementation, as long as the programmer also ensures that the code is idiomatic.

 **What is your recommendation to the firm regarding its staffing of programmers and software engineers?**
 <br>
 I would recommend that the firm maintain its current levels of programmers and software engineers, but allow them to use AI-Assisted programming. This will ensure that the work is done faster, but also that knowledge is retained in the workforce. I would not recommend AI-generated code, because human programming is more proficient (at least for Go), particularly with identifying relevant packages and resources.

 ## Use of AI
For this assignment, I used [Github Copilot in VSCode](https://code.visualstudio.com/docs/copilot/overview) for the AI-Assisted portion and [ChatGPT](https://chatgpt.com/) for the AI-Generated Code portion. The Automated Code Generation section, I attempted to use Github Copilot's Code Completions in VSCode, but they were not very effective and I ended up turning them off. 

For more information, see the Assignment_5 Notes.txt document in this repository.

### Automated Code Generation
The Automated Code Generation provided by Copilot when coding in VSCode was not effective. It attempted to auto-complete based on its knowledge-base and documents it could access from the directory. In actuality, it was a more intrusive auto-complete mechanism that could not accurately guess what I was trying to complete. I would recommend turning this off in an enterprise solution.

### AI-Assisted Programming
For AI-Assisted programming, I used the Github Copilot in VSCode to make adjustments and edits to the code that I had already created. Initially, the Chatbot calculated the linear regression coefficients without using a package, so I had to specify which package I wanted it to use. I have mentioned before that I liked the test cases it produced, but there did reach a point where it was testing code that could be addressed in the functions themselves. Copilot's tendency to favor for-loops was especially impactful when generating code for Python and R. In these cases, the code ran much slower so I had to specify to relax on the for-loops. It did produce more complete code immediately after I asked for Python and R. I was able to coax it to produce idiomatic Go code, but I had to specify that is what I wanted.
### AI Generated Code
- Step-by-Step Review of Work
- Prompt review
- Describe how files were obtained

## Background
The purpose of the Anscombe Quartet is to demonstrate how it is important to visualize datasets instead of relying on pure data. Each of the datasets in the quartet returns identifical regression coefficients, despite being very different. 

For more information, see:
Anscombe, F. J. 1973, February. "Graphs in Statistical Analysis." *The American Statistician 27(1)*: 17–21. Available online at [https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf).


## Roles of Programs and Data
Note any training materials.

### AI_Assisted
- go.mod:
- go.sum:
- Linear Regression with R.R:
- main_test.go:
- main.go:
- Python File for Assignment 4.py:

### AI_Generated
- go.mod:
- go.sum:
- Linear Regression with R.R:
- main_test.go:
- main.go:
- Python File for Assignment 4.py:

## Application
text


