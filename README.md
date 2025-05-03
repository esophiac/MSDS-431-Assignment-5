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

The full log for completing this assignment with Github Copilot can be found in Week 5 Assignment VSCode Log.txt.

### AI Generated Code
The first thing I did with ChatGPT was ask if it was familiar with the Anscombe Quartet. It said it was, so I asked it to generate the full Go code for calculating the linear regression coefficients. It also didn't initially use any packages besides fmt, and also got the Anscombe Quartet wrong in the body of the code. When I asked it to fix it, it moved everything out of functions and into the body of the Go code. I then worked on making the code more idiomatic. It performed similiarly with R and Python. I might have gotten better results if I had specified specific packages that I know are effective, but for the sake of this test I was going to run the programs that ChatGPT gave me as is. The only time I specified a package/library was to give it the same instructions for Assignment 4.

The full log for completing this assignment with ChatGPT can be found in Week 5 Assignment ChatGPT Log.txt.

## Background
This assignment tested how AI-assisted programming performs in comparison to the task set in Assignment 4; calculate the linear regression coefficients of the Anscombe Quartet using the [stats package](https://pkg.go.dev/github.com/montanaflynn/stats).

The purpose of the Anscombe Quartet is to demonstrate how it is important to visualize datasets instead of relying on pure data. Each of the datasets in the quartet returns identifical regression coefficients, despite being very different. 

For more information, see:
Anscombe, F. J. 1973, February. "Graphs in Statistical Analysis." *The American Statistician 27(1)*: 17–21. Available online at [https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf](https://www.sjsu.edu/faculty/gerstman/StatPrimer/anscombe1973.pdf).


## Roles of Programs and Data
No training materials were used in this assignment. To make sure that the Chatbots had access to the Anscombe Quartet, I first asked them if they did. I had to provide a small correction to ChatGPT, but that was it.

To make it easier to organize the project, I created sub-directories for AI_Assisted and AI_Generated, which house the files relevant to each test. Other files, including logs for discussions with the chat agents, are also in the repository.

- Assignment_5 Notes.txt: additional notes and observations relevant to the tests
- Week 5 Assignment ChatGPT Log.txt: the log of the conversation with ChatGPT for the AI Generated test
- Week 5 Assignment VSCode Log.txt: the log of the conversation with Github Copilot for the AI Assisted test
- README.md: the readme file for the repository

### AI_Assisted
- Assignment_5.exe: the executable for this part of the project. Assembled on Windows.
- go.mod:defines the module's properties (for the AI_Assisted Go files)
- go.sum: record of the library the AI_Assisted project depends on
- Linear Regression with R.R: the file to test the Anscombe quartet with the R programming language. The standard library in R was used.
- main_test.go: tests and benchmarks the fuctions in the main.go file in this directory
- main.go: the file to test the Anscombe quartet with the Go programming language with AI assistance. The stats library was used.
- Python File for Assignment 4.py: the file to test the Anscombe quartet with the Python programming language. Used the same packages as Assignment 4. 

### AI_Generated
- AI_generated.exe: the executable for this part of the project. Assembled on Windows.
- go.mod: defines the module's properties (for the AI_Generated Go files)
- go.sum: record of the library the AI_Generated project depends on
- Linear Regression with R.R: the file to test the Anscombe quartet with the R programming language. The standard library in R was used. ChatGPT also pulled the Anscombe Quartet from the standard data that R provides instead of defining it itself. 
- main_test.go: tests and benchmarks the fuctions in the main.go file in this directory
- main.go: the file to test the Anscombe quartet with the Go programming language with AI-generated code. The stats library was used.
- Python File for Assignment 4.py: the file to test the Anscombe quartet with the Python programming language. Only used numpy for the calculations.

## Application
An executable for each part of this project was created using Windows and is included in the subfolders in this repository. To create your own executable, run **go build** in the same directory as the go program. For more information, see the Gopher documentation on creating an executable [here](https://go.dev/doc/tutorial/compile-install).

