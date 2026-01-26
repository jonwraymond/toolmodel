 #!/bin/bash
set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <iterations> [prd_file] [progress_file]"
  echo "  iterations:   number of iterations to run"
  echo "  prd_file:     path to PRD file (default: PRD.md)"
  echo "  progress_file: path to progress file (default: progress.txt)"
  exit 1
fi

ITERATIONS=$1
PRD_FILE=${2:-PRD.md}
PROGRESS_FILE=${3:-progress.txt}

for ((i=1; i<=$ITERATIONS; i++)); do
  # Use local claude instead of docker sandbox due to tool use bug in claude 2.1.19
  result=$(claude --permission-mode acceptEdits -p "@$PRD_FILE @$PROGRESS_FILE \
  1. Find the highest-priority task and implement it. \
  2. Run your tests and type checks. \
  3. Update the PRD with what was done. \
  4. Append your progress to $PROGRESS_FILE. \
  5. Commit your changes. \
  ONLY WORK ON A SINGLE TASK. \
  If the PRD is complete, output <promise>COMPLETE</promise>.")

  echo "$result"

  if [[ "$result" == *"<promise>COMPLETE</promise>"* ]]; then
    echo "PRD complete after $i iterations."
    exit 0
  fi
done
