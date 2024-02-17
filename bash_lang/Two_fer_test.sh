#!/usr/bin/env bats
load bats-extra

@test "no name given" {
  run bash Two_fer.sh    
  assert_success
  assert_output "One for you, one for me."
}

@test "a name given" {
  [[ $BATS_RUN_SKIPPED == "true" ]] || skip
  run bash Two_fer.sh Alice
  assert_success
  assert_output "One for Alice, one for me."
}

@test "another name given" {
  [[ $BATS_RUN_SKIPPED == "true" ]] || skip
  run bash Two_fer.sh Bob
  assert_success
  assert_output "One for Bob, one for me."
}

@test "handle arg with spaces" {
  [[ $BATS_RUN_SKIPPED == "true" ]] || skip
  run bash Two_fer.sh "John Smith" "Mary Ann"
  assert_success
  assert_output "One for John Smith, one for me."
}          

@test "handle arg with glob char" {
  [[ $BATS_RUN_SKIPPED == "true" ]] || skip
  run bash Two_fer.sh "* "
  assert_success
  assert_output "One for * , one for me."
}
