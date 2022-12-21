package com.platzi.jobsearch.cli;

import com.beust.jcommander.Parameter;
import lombok.Data;


@Data
public class CliArguments {
    @Parameter(
            required = true,
            descriptionKey = "KEYWORD",
            validateWith = CliKeywordValidator.class,
            description = "KEYWORD"
    )
    private String keyword;
    @Parameter(
            names = {"--location", "-l"},
            description = "Busqueda correspondiente por pais y ciudad"
    )
    private String location;
    @Parameter(
            names = {"--page", "-p"},
            description = "Pagina donde nos paramos por temas de paginación"
    )
    private int page = 0;
    @Parameter(
            names = {"--full-time"},
            description = "Agregar empleos de tiempo completo"
    )
    private boolean isFullTime = false;
    @Parameter(
            names = {"--markdown"},
            description = "Obtener los resultados en markdown"
    )
    private boolean isMarkdown = false;
    @Parameter(
            names = {"--help"},
            help = true,
            validateWith = CliHelpValidator.class,
            description = "Monstrar esta ayuda"
    )
    private boolean isHelp;

    public static CliArguments newInstance() {
        return new CliArguments();
    }
}
