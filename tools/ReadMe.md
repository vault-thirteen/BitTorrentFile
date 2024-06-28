# Tools

Auxiliary tools for _BitTorrent_ files.

* Indexer

## Indexer

The `Indexer` tool collects information about files hashed in _BitTorrent_ 
files and saves this information into a file of _CSV_ format.

The indexer tool can index a single _BitTorrent_ file or a whole folder 
containing any number of _BitTorrent_ files. If a folder contains files of 
various types (i.e. files having different extensions), the tool skips all 
files which do not have the `torrent` file extension. The tool reads folder's 
content recursively, i.e. it is able to read files sitting inside subfolders of 
a folder.

To see how to use the tool, start it without command line arguments.
