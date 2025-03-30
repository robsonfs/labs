#!/bin/bash

# Input video file
INPUT_VIDEO="input.mp4"
SCENE_THRESHOLD=0.4  # Adjust for sensitivity (0.3 to 0.5 is a good range)
OUTPUT_DIR="scenes"

mkdir -p "$OUTPUT_DIR"

# Step 1: Extract scene timestamps using ffprobe
ffprobe -hide_banner -i "$INPUT_VIDEO" -vf "select='gt(scene,$SCENE_THRESHOLD)',showinfo" \
    -show_entries frame=pkt_pts_time -of csv=p=0 2>&1 | grep "pts_time" | awk -F'=' '{print $2}' > scene_timestamps.txt

echo "Detected scene timestamps:"
cat scene_timestamps.txt

# Step 2: Split video using extracted timestamps
START_TIME=0
INDEX=1

timestamps=($(cat scene_timestamps.txt))

echo "Splitting video into scenes..."
for TIMESTAMP in "${timestamps[@]}"; do
    OUTPUT_FILE="$OUTPUT_DIR/scene_$(printf "%03d" $INDEX).mp4"
    echo "Extracting scene $INDEX: Start=$START_TIME, End=$TIMESTAMP"
    ffmpeg -i "$INPUT_VIDEO" -ss "$START_TIME" -to "$TIMESTAMP" -c copy "$OUTPUT_FILE"
    START_TIME=$TIMESTAMP
    INDEX=$((INDEX + 1))
done

# Process the last segment (remaining part of the video)
OUTPUT_FILE="$OUTPUT_DIR/scene_$(printf "%03d" $INDEX).mp4"
echo "Extracting last scene: Start=$START_TIME, until end of video"
ffmpeg -i "$INPUT_VIDEO" -ss "$START_TIME" -c copy "$OUTPUT_FILE"

echo "Scene splitting complete. Check the '$OUTPUT_DIR' folder."
