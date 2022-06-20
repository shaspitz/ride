#!/bin/bash
unset bob
unset alice
export alice=$(rided keys show alice -a)
export bob=$(rided keys show bob -a)
rided tx ride finish 1 "some loc" --from bob --gas auto