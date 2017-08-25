package lagou

import (
	"encoding/json"
	"fmt"

	"github.com/WindomZ/lagou-jobs/lagou/entity/mobile"
	"github.com/WindomZ/lagou-jobs/lagou/output"
)

func (s Spider) initFiles() error {
	if len(s.Config.Output.Files.JSON) != 0 {
		if err := output.Access(s.Config.Output.Files.JSON); err != nil {
			return err
		}
	}
	return nil
}

func (s Spider) writeToFiles(p mobile.PositionMap) error {
	if len(s.Config.Output.Files.JSON) != 0 { // output json file
		r := Results{}
		if err := r.fromPositionMap(p); err != nil {
			return err
		}

		if b, err := json.Marshal(r); err != nil {
			return err
		} else if err := output.Write(s.Config.Output.Files.JSON,
			string(b)); err != nil {
			return err
		}

		fmt.Printf("Success output JSON file: '%s'\n", s.Config.Output.Files.JSON)
	}
	return nil
}
