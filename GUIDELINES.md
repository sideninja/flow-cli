# CLI Style Guide
This is a draft guideline meant for development of Flow CLI. Idea is to have certain rules developers of CLI should follow for users to have a feeling of consistency and logical usage.

## Naming
Commands should be singular nouns in lower case without whitespaces and should use dash (-) for multi word commands. Command should be followed by a verb. Command names should be clear and follow some established command names seen in other CLIs. Commands such as help and version must be included in the set. 
Naming and language should be same as used in Flow documentation.

```
flow account
```

## Flags
Flags act as input and can be represented in longer and shorter versions. Example --username or shorter -u. Longer version is preferred for better readability and simpler learning experience in the beginning but after users will probably start using the shorter version. Every flag should have a shorter version and they should be both presented in the help. Every flag that can have a default value should have a default value so users don’t have to specify all of them when not needed. Support flags following a value delimited by a space or equal (=) sign (ex: --username test or --username=test)

```
flow account get --filter "address"
```

## Arguments
Arguments are meant to be used to specify a subcommand. There should not be more than one argument because it makes it hard for users to remember the sequence of arguments. Because they are relaying on position flags should be used where more than one argument would be required. Arguments should be one word verbs following general naming guideline.

```
flow account get
```

## Interaction
Commands must be stateless, meaning they can be run without relying on external state. If we have some external state like url config each command must include an option to define it on the fly, but it must also work without by first relaying on the externally saved state in config and if not found relaying on default value. If when relaying on default value the value creates an error there should be an explanation for the user that config must be created. We try to use default values first to get that “works like magic” feeling.


## Help
Main help screen must contain a list of all commands with their own description, it must also include examples for commands.
Commands should have a description that is 80 characters long or less. Help should be outputed if command is ran without any arguments. Help outline should be:

```Description:
    <description>

Usage:
    <command> <action>
    <command> <action>
    …

Examples:
    An optional section of example(s) on how to run the command.

Commands:
    <command>
        <commandDescription>
    <command>
        <commandDescription>
    …

Options:
    --<flag>
        <flagDescription>
    --<flag>
        <flagDescription>
    …
```

## Progress
Show progress for longer running commands visually. Visual feedback on longer running commands is important so users don’t get confused if command is running or the CLI is hang up resulting in user quitting. 

## Feedback
Commands should provide feedback if there is no result to be presented. If a command completes an operation that has no result to be shown it should write out that command was executed. This is meant to assure the user of the completion of the command.
If user is executing a destructive command they should be asked for approval but this command should allow passing confirmation with a flag --yes. 

## Exit
CLI should return zero status unless it is shut down because of an error. This will allow chaining of commands and interfacing with other cli. If a command is long running then provide description how to exit it (press ctrl+c to stop).
```
exit status 1
```

## Error
Error should be human readable, avoid printing out stack trace. It should give some pointers as to what could have gone wrong and how to fix it. Maybe try involving an example. However allow --verbose flag for complete error info with stack.

```
Connection to Flow Emulator was not successful, please make sure your api url is correct.

Set it by using: flow config 

Connection Error: connection error: desc = "transport: Error while dialing dial tcp [::1]:3000: connect: connection refused"
```

## Colors
Base color should be white with yellow reserved for warnings and red for errors. Blue color can be used to accent some commands but it shouldn’t be used too much as it will confuse the user.

## Output
Output of commands should be presented in a clear formatted way for users to easily scan it, but it must also be compatible with grep command often used in command chaining. Output should also be possible in json by using --json flag. 

```
Address  179b6b1cb6755e31
Balance  0
Keys     2

Key 0   Public Key               c8a2a318b9099cc6c872a0ec3dcd9f59d17837e4ffd6cd8a1f913ddfa769559605e1ad6ad603ebb511f5a6c8125f863abc2e9c600216edaa07104a0fe320dba7
        Weight                   1000
        Signature Algorithm      ECDSA_P256
        Hash Algorithm           SHA3_256


Key 1   Public Key               0a4c9f7bf0c8adf6ebdf4859c11d8e2867e0eaa4af6880e18a0730a0b44a494e976cefa0caf8efb7ec2da469c3f320dab4a2ca72fb340621776f4a1403ae39ed
        Weight                   1000
        Signature Algorithm      ECDSA_P256
        Hash Algorithm           SHA3_256


Code             
         pub contract Foo {
                pub var bar: String
         
                init() {
                        self.bar = "Hello, World!"
                }
         }
```

```
{"Address":"179b6b1cb6755e31","Balance":0,"Code":"CnB1YiBjb250cmFjdCBGb28gewoJcHViIHZhciBiYXI6IFN0cmluZwoKCWluaXQoKSB7CgkJc2VsZi5iYXIgPSAiSGVsbG8sIFdvcmxkISIKCX0KfQo=","Keys":[{"Index":0,"PublicKey":{},"SigAlgo":2,"HashAlgo":3,"Weight":1000,"SequenceNumber":0,"Revoked":false},{"Index":1,"PublicKey":{},"SigAlgo":2,"HashAlgo":3,"Weight":1000,"SequenceNumber":0,"Revoked":false}],"Contracts":null}
```






# Inspiration and Reference 
https://blog.developer.atlassian.com/10-design-principles-for-delightful-clis/

https://devcenter.heroku.com/articles/cli-style-guide

https://eng.localytics.com/exploring-cli-best-practices/








