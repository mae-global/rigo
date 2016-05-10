package rigo

import (
	. "github.com/mae-global/rigo/ri"
)

/* These _pipes_ are for fixing known issues. Do NOT include them in your own pipes as they will
 * be automagically added as required.
 */

/* https://github.com/mae-global/rigo/issues/1 */

type PipeIssue0001Fix struct {
	Insession   bool /* have you seen a RiBegin */
	VersionSeen bool /* have you already seen a version name token*/
}

func (p *PipeIssue0001Fix) ToRaw() ArchiveWriter { return nil }
func (p *PipeIssue0001Fix) Name() string         { return "pipe-issue-0001-fix" }
func (p *PipeIssue0001Fix) Pipe(name RtName, args, params, values []Rter, info Info) *Result {

	switch string(name) {
	case "Begin":
		p.Insession = true
		p.VersionSeen = false
		break
	case "version":

		if p.VersionSeen || !p.Insession {
			return Skip()
		} else {
			p.VersionSeen = true
		}
		break
	case "End":
		p.Insession = false
		p.VersionSeen = false
		break
	}

	return Done()
}
