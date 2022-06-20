#!/bin/bash
unset bob
unset alice
export alice=$(rided keys show alice -a)
export bob=$(rided keys show bob -a)
rided tx ride request-ride "some dest" "some other dest" 50 5 25 --from alice --gas auto