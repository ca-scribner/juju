# FAQ


Q: I want to get involved.  Where should I start?
A: Read the docs.  They're not complete (feel free to contribute), but
   they help.  Suggested reading order:

     1. README
     2. CONTRIBUTING
     3. docs/architectural-overview.txt
     4. docs/how-to-write-tests.txt
     5. the other docs in docs/

   Many of the docs in the docs/ directory become progressively more clear 
   as you begin actually working in the code.  Keep re-reading them
   periodically to gain more benefit.



Q: How do I find out what's going on currently?  How does the juju-core 
   team communicate regularly?
A: IRC.  Mailing Lists.  Code Reviews.

     1. IRC: most day-to-day discussion happens in #juju on freenode.
     2. Mailing Lists: [juju-dev](https://lists.ubuntu.com/mailman/listinfo/juju-dev)
     3. Code Reviews: using github [Pull Requests](https://github.com/juju/juju/pulls)



Q: How do I run only specific test suites using gocheck?
A: The quick answer is that the -gocheck parameters must be *last* on the
   command line, for example:

     go test -v ./state -gocheck.v -gocheck.f="StateSuite|CleanupSuite"

   Also, if you want to use -gocheck.v or -gocheck.vv, you must also provide 
   the -v option to go test, otherwise the -gocheck.v* options are silently
   ignored.

   Additionally, you can only specify -gocheck.f filters on one package at
   a time.  If you use the go test ellipsis (./...) it will ignore your
   filter strings.



