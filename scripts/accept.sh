#!/bin/bash
unset bob
unset alice
export alice=$(rided keys show alice -a)
export bob=$(rided keys show bob -a)
rided tx ride accept 1 --from bob --gas auto