package com.platzi.functional._04_functional;

import java.util.function.Consumer;
import java.util.function.Supplier;

public class CLIArgumentUtils {
    static void showHelp(CLIArgument cliArgument) {
        Consumer<CLIArgument> consumerHelper = cliArgument1 -> {
            if (cliArgument.isHelp()) {
                System.out.println("Manual solicitado");
            }
        };

        consumerHelper.accept(cliArgument);
    }

    static CLIArgument generateCLI(){
        Supplier<CLIArgument> generator = () -> new CLIArgument();
        return generator.get();
    }
}
