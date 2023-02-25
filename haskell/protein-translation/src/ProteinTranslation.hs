module ProteinTranslation(proteins) where

proteins :: String -> Maybe [String]
proteins s = Just (codons s)

codons :: String -> [String]
codons [] = []
codons s = build (protein (take 3 s)) (drop 3 s)

build :: String -> String -> [String]
build "STOP" _ = []
build s rest = s : codons rest

protein :: String -> String
protein s = case s of
    "AUG" -> "Methionine"
    "UUU" -> "Phenylalanine"
    "UUC" -> "Phenylalanine"
    "UUA" -> "Leucine"
    "UUG" -> "Leucine"
    "UCU" -> "Serine"
    "UCC" -> "Serine"
    "UCA" -> "Serine"
    "UCG" -> "Serine"
    "UAU" -> "Tyrosine"
    "UAC" -> "Tyrosine"
    "UGU" -> "Cysteine"
    "UGC" -> "Cysteine"
    "UGG" -> "Tryptophan"
    "UAA" -> "STOP"
    "UAG" -> "STOP"
    "UGA" -> "STOP"
    _     -> error "Invalid codon"
