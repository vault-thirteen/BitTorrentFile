package main

const (
	MsgFInspectedBTFilesCount = "Inspected BitTorrent files: %d."
	MsgFProcessingFiles       = "Processing %d BitTorrent files:"
	MsgProcessingHasFinished  = "Processing has finished."
	MsgFTimeTakenSeconds      = "Time taken: %f seconds."
	MsgFBytesRead             = "Bytes read: %d."
	MsgFBytesWritten          = "Bytes written: %d."
	MsgFFilesIndexed          = "Files indexed: %d."
	MsgFFilesBroken           = "Broken files: %d."
	MsgFSelfCheckErrors       = "Self-check errors: %d."
)

const UsageHint = `Usage:
	[ObjectType] [ObjectPath] [Output]

Examples:
	indexer.exe File "123.torrent" "index.csv" 
	indexer.exe Folder "torrents" "index.csv"

Notes:
	Possible object types: File, Folder or Directory.
	Letter case is not important.
	Change directory (CD) to a working directory before usage.`

const (
	ErrBrokenFile = "broken file"
)
