#!/bin/sh
echo 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 > .token && gh auth login --with-token < "./.token" && rm ./.token
	
gh repo create "jackkweyunga/automatedrepo" --private --confirm