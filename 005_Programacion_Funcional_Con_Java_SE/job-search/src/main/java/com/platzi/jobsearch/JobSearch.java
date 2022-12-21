package com.platzi.jobsearch;

import com.beust.jcommander.JCommander;
import com.platzi.jobsearch.api.IApiJobs;
import com.platzi.jobsearch.cli.CliArguments;
import com.platzi.jobsearch.cli.CliFunctions;

import java.util.Collection;
import java.util.Collections;
import java.util.Map;
import java.util.Optional;
import java.util.stream.Stream;

import static com.platzi.jobsearch.CommanderFunctions.buildCommanderWithName;
import static com.platzi.jobsearch.CommanderFunctions.parseArguments;
import static com.platzi.jobsearch.api.IApiFunctions.buildApi;

public class JobSearch {
    public static void main(String[] args) {
        JCommander jCommander = buildCommanderWithName("job-search", CliArguments::newInstance);

        Stream<CliArguments> streamOfCli = parseArguments(jCommander, args, JCommander::usage)
                .orElse(Collections.emptyList())
                .stream()
                .map(o -> (CliArguments) o);

        Optional<CliArguments> cliArgumentsOptional = streamOfCli.filter(cliArguments -> !cliArguments.isHelp())
                .filter(cliArguments -> cliArguments.getKeyword() != null)
                .findFirst();

        cliArgumentsOptional.map(CliFunctions::toMap)
                .map(JobSearch::executeRequest)
                .orElse(Stream.empty())
                .forEach(System.out::println);
    }

    private static Stream<JobPosition> executeRequest(Map<String, Object> params) {
        IApiJobs apiJobs = buildApi(IApiJobs.class, "https://jobs.github.com");
        return Stream.of(params)
                .map(apiJobs::jobsAvailable)
                .flatMap(Collection::stream);
    }
}
