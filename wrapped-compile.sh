#!/bin/bash


project_path=$1
current_path=$(pwd)

cd $project_path
cargo build-bpf &> ./output.log
cd $current_path
