#!/bin/bash
unset bob
unset alice
export alice=$(rided keys show alice -a)
export bob=$(rided keys show bob -a)
rided tx ride rate 1 $bob 8.5 --from alice --gas auto 