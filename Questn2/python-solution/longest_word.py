def longest_word(sentence):
  words = sentence.split()
  wordMapLength = dict()
  maxLen = len(words[0])
  maxStr = words[0]
  for word in words:
    wordMapLength[word] = len(word)
    if len(word) >= maxLen:
      maxLen = len(word)
      maxStr = word

  return maxStr