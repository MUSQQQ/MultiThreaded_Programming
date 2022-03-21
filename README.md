# README

Small repo for my multithreaded projects from uni. Mainly for GoLang i guess.

# 1. Integral calculations
I program - wrong calculation of an integral using goroutines (without any control over access to the main variable)
II program - calculation of an integral using goroutines and mutexes
III program - calculation of an integral using goroutines and critical section (and mutexes)
IV program - calculation of an integral using a different method than those mentioned above. Right now it's just by using channels.

# 2. Dining philosophers
GoLang implementation of the Dining Philosophers Problem with a twist. Philosophers randomly choose their seats and they take forks from both sides of the chosen seats. They also choose one book that they are going to read while eating.

# 3. Seqlock implementation
GoLang implementation of the Linux kernels seqlock. Simplified.
