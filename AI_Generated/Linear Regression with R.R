

# Load Anscombe dataset
data(anscombe)

# Start time
start_time <- Sys.time()

# Function to compute slope and intercept
compute_regression <- function(x, y) {
  model <- lm(y ~ x)
  coef(model)
}

# Loop through datasets
for (i in 1:4) {
  x <- anscombe[[paste0("x", i)]]
  y <- anscombe[[paste0("y", i)]]
  
  coefficients <- compute_regression(x, y)
  cat(sprintf("Dataset %d: Slope = %.2f, Intercept = %.2f\n",
              i, coefficients["x"], coefficients["(Intercept)"]))
}

# End time and duration
end_time <- Sys.time()
duration <- end_time - start_time
cat(sprintf("Execution time: %.6f seconds\n", as.numeric(duration, units = "secs")))
