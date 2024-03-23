Giving lsb stenography a try.

Here is what I hope this will do in no particular order:

1. Take an image as input and generate a header that states the total number of possible entries based on image resolution.
2. In the same header, store how many entries are already stored in the image.
3. Take an image name and key-value string pair as input and store that data in an image. This will also update existing values if the key already exists. (possibly -s --store for this)
4. Take an image as input and a key only to look up if that key exists in the image. If it does return the key-value pair. (possibly -g --get for this)
5. Take an image as input and clear all lsb's. Use the --microwave flag for this.


Ideas for each section:

1.
2. Just increment or decrement this based on whether new entries have been added.
3. I can just loop through chunks of sizeof(key) lsb's in the image until I have a matching result? Or I can limit the key size to 128 bytes. This would make it easy to iterate over the keys. We would also need some sort of metadata that describes the value length so we can make it as large as we need. But that might make it easy for someone else to search for the keys and grab the data. This makes it not so secret. But maybe we don't care that much! The first idea makes it so that only the person who wrote the key in to the file knows what keys are in the file. Maybe we can implement both ideas using a different flag for each. (--structured vs. --unstructured storage maybe?)
4.
5.
